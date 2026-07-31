package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	URL         string
	AppName     string
	PackageName string
	OutputDir   string
	IconPath    string
	BuildDir    string
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.URL, "url", "", "URL Website yang akan diubah menjadi APK")
	flag.StringVar(&cfg.AppName, "name", "MyWebApp", "Nama Aplikasi Android")
	flag.StringVar(&cfg.PackageName, "package", "com.mycompany.webapp", "Package Name Android")
	flag.StringVar(&cfg.OutputDir, "out", ".", "Folder output APK")
	flag.StringVar(&cfg.IconPath, "icon", "", "Path ke file gambar ikon aplikasi (.png / .jpg)")
	flag.Parse()

	if cfg.URL == "" {
		fmt.Println("\n❌ Error: Flag -url wajib diisi!")
		fmt.Println("Penggunaan: web2apk -url https://example.com -name MyWebsite")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !strings.HasPrefix(cfg.URL, "http://") && !strings.HasPrefix(cfg.URL, "https://") {
		cfg.URL = "https://" + cfg.URL
	}

	u, err := url.Parse(cfg.URL)
	if err != nil || u.Host == "" {
		fmt.Printf("❌ Format URL '%s' tidak valid!\n", cfg.URL)
		os.Exit(1)
	}

	fmt.Println("==================================================")
	fmt.Println("       🚀 GO WEB2APK BUILDER TOOL 🚀")
	fmt.Println("==================================================")
	fmt.Printf("🌐 Target URL    : %s\n", cfg.URL)
	fmt.Printf("📱 App Name      : %s\n", cfg.AppName)
	fmt.Printf("📦 Package Name  : %s\n", cfg.PackageName)
	fmt.Printf("📂 Output Folder : %s\n", cfg.OutputDir)
	fmt.Println("==================================================")

	tmpDir, err := os.MkdirTemp("", "web2apk-build-*")
	if err != nil {
		fmt.Printf("❌ Gagal membuat direktori build sementara: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)
	cfg.BuildDir = tmpDir

	fmt.Println("🛠️  [1/6] Menyiapkan struktur projek Android...")
	if err := generateSourceFiles(cfg); err != nil {
		fmt.Printf("❌ Gagal membuat source code Android: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("⚙️  [2/6] Memproses Resource & Manifest dengan AAPT...")
	if err := runAAPT(cfg); err != nil {
		fmt.Printf("❌ AAPT Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("☕ [3/6] Mengompilasi Java Source Code (ECJ)...")
	if err := compileJava(cfg); err != nil {
		fmt.Printf("❌ Java Compilation Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("⚡ [4/6] Mengonversi Class ke Dalvik Executable (DX/DEX)...")
	if err := runDX(cfg); err != nil {
		fmt.Printf("❌ DEX Conversion Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📦 [5/6] Mengemas Unsigned APK...")
	unsignedAPK := filepath.Join(cfg.BuildDir, "app-unsigned.apk")
	if err := buildUnsignedAPK(cfg, unsignedAPK); err != nil {
		fmt.Printf("❌ Gagal mengemas APK: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🔑 [6/6] Melakukan Signing APK (Keystore & APKSigner)...")
	finalAPKName := fmt.Sprintf("%s.apk", strings.ToLower(strings.ReplaceAll(cfg.AppName, " ", "_")))
	finalAPKPath := filepath.Join(cfg.OutputDir, finalAPKName)

	if err := signAPK(cfg, unsignedAPK, finalAPKPath); err != nil {
		fmt.Printf("❌ APK Signing Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("==================================================")
	fmt.Println("🎉 BERHASIL! APK telah selesai dibuat.")
	fmt.Printf("📌 Lokasi File: %s\n", finalAPKPath)
	fmt.Println("==================================================")
}

func generateSourceFiles(cfg Config) error {
	pkgPath := strings.ReplaceAll(cfg.PackageName, ".", "/")
	javaSrcDir := filepath.Join(cfg.BuildDir, "src", pkgPath)
	resDir := filepath.Join(cfg.BuildDir, "res", "layout")
	valuesDir := filepath.Join(cfg.BuildDir, "res", "values")
	mipmapDir := filepath.Join(cfg.BuildDir, "res", "mipmap-hdpi")

	if err := os.MkdirAll(javaSrcDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(resDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(valuesDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(mipmapDir, 0755); err != nil {
		return err
	}

	manifestContent := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="%s"
    android:versionCode="1"
    android:versionName="1.0">

    <uses-sdk android:minSdkVersion="24" android:targetSdkVersion="34" />

    <uses-permission android:name="android.permission.INTERNET" />
    <uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
    <uses-permission android:name="android.permission.READ_MEDIA_IMAGES" />
    <uses-permission android:name="android.permission.READ_MEDIA_VIDEO" />
    <uses-permission android:name="android.permission.READ_MEDIA_AUDIO" />
    <uses-permission android:name="android.permission.POST_NOTIFICATIONS" />
    <uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE" android:maxSdkVersion="32" />
    <uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE" android:maxSdkVersion="32" />

    <application
        android:label="@string/app_name"
        android:icon="@mipmap/ic_launcher"
        android:usesCleartextTraffic="true"
        android:hardwareAccelerated="true"
        android:enableOnBackInvokedCallback="true">
        <activity
            android:name=".MainActivity"
            android:configChanges="orientation|screenSize|keyboardHidden"
            android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>
</manifest>`, cfg.PackageName)

	if err := os.WriteFile(filepath.Join(cfg.BuildDir, "AndroidManifest.xml"), []byte(manifestContent), 0644); err != nil {
		return err
	}

	stringsContent := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="app_name">%s</string>
</resources>`, cfg.AppName)
	if err := os.WriteFile(filepath.Join(valuesDir, "strings.xml"), []byte(stringsContent), 0644); err != nil {
		return err
	}

	layoutContent := `<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android">
</LinearLayout>`
	if err := os.WriteFile(filepath.Join(resDir, "activity_main.xml"), []byte(layoutContent), 0644); err != nil {
		return err
	}

	iconDest := filepath.Join(mipmapDir, "ic_launcher.png")
	if cfg.IconPath != "" && fileExists(cfg.IconPath) {
		ext := strings.ToLower(filepath.Ext(cfg.IconPath))
		if ext == ".jpg" || ext == ".jpeg" {
			cmdConv := exec.Command("magick", cfg.IconPath, iconDest)
			if err := cmdConv.Run(); err != nil {
				cmdConv2 := exec.Command("convert", cfg.IconPath, iconDest)
				if err2 := cmdConv2.Run(); err2 != nil {
					return fmt.Errorf("gagal mengonversi file ikon JPG ke PNG: %v", err2)
				}
			}
		} else {
			if err := copyFile(cfg.IconPath, iconDest); err != nil {
				return fmt.Errorf("gagal menyalin file ikon: %v", err)
			}
		}
	} else {
		defaultPng := "/data/data/com.termux/files/usr/share/doc/libogg/fish_xiph_org.png"
		if fileExists(defaultPng) {
			_ = copyFile(defaultPng, iconDest)
		} else {
			_ = os.WriteFile(iconDest, []byte{}, 0644)
		}
	}

	javaContent := fmt.Sprintf(`package %s;

import android.app.Activity;
import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;
import android.webkit.ValueCallback;
import android.webkit.WebChromeClient;
import android.webkit.WebSettings;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.view.KeyEvent;

public class MainActivity extends Activity {
    private WebView mWebView;
    private ValueCallback<Uri[]> mFilePathCallback;
    private static final int FILECHOOSER_RESULTCODE = 1;
    private static final String TARGET_URL = "%s";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        
        mWebView = new WebView(this);
        setContentView(mWebView);

        WebSettings webSettings = mWebView.getSettings();
        webSettings.setJavaScriptEnabled(true);
        webSettings.setDomStorageEnabled(true);
        webSettings.setDatabaseEnabled(true);
        webSettings.setAllowFileAccess(true);
        webSettings.setAllowContentAccess(true);
        webSettings.setUseWideViewPort(true);
        webSettings.setLoadWithOverviewMode(true);
        webSettings.setBuiltInZoomControls(false);
        webSettings.setMediaPlaybackRequiresUserGesture(false);
        webSettings.setMixedContentMode(WebSettings.MIXED_CONTENT_ALWAYS_ALLOW);

        mWebView.setWebViewClient(new CustomWebViewClient());
        mWebView.setWebChromeClient(new CustomWebChromeClient());

        mWebView.loadUrl(TARGET_URL);
    }

    private class CustomWebViewClient extends WebViewClient {
        @Override
        public boolean shouldOverrideUrlLoading(WebView view, String url) {
            if (url.startsWith("http://") || url.startsWith("https://")) {
                view.loadUrl(url);
                return true;
            }
            try {
                Intent intent = new Intent(Intent.ACTION_VIEW, Uri.parse(url));
                startActivity(intent);
                return true;
            } catch (Exception e) {
                return false;
            }
        }
    }

    private class CustomWebChromeClient extends WebChromeClient {
        @Override
        public boolean onShowFileChooser(WebView webView, ValueCallback<Uri[]> filePathCallback, FileChooserParams fileChooserParams) {
            if (mFilePathCallback != null) {
                mFilePathCallback.onReceiveValue(null);
            }
            mFilePathCallback = filePathCallback;

            Intent intent = fileChooserParams.createIntent();
            try {
                startActivityForResult(intent, FILECHOOSER_RESULTCODE);
            } catch (Exception e) {
                mFilePathCallback = null;
                return false;
            }
            return true;
        }
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        if (requestCode == FILECHOOSER_RESULTCODE) {
            if (mFilePathCallback == null) return;
            Uri[] results = null;
            if (resultCode == RESULT_OK && data != null) {
                String dataString = data.getDataString();
                if (dataString != null) {
                    results = new Uri[]{Uri.parse(dataString)};
                } else if (data.getClipData() != null) {
                    int count = data.getClipData().getItemCount();
                    results = new Uri[count];
                    for (int i = 0; i < count; i++) {
                        results[i] = data.getClipData().getItemAt(i).getUri();
                    }
                }
            }
            mFilePathCallback.onReceiveValue(results);
            mFilePathCallback = null;
        } else {
            super.onActivityResult(requestCode, resultCode, data);
        }
    }

    @Override
    public boolean onKeyDown(int keyCode, KeyEvent event) {
        if ((keyCode == KeyEvent.KEYCODE_BACK) && mWebView.canGoBack()) {
            mWebView.goBack();
            return true;
        }
        return super.onKeyDown(keyCode, event);
    }
}`, cfg.PackageName, cfg.URL)

	return os.WriteFile(filepath.Join(javaSrcDir, "MainActivity.java"), []byte(javaContent), 0644)
}

func runAAPT(cfg Config) error {
	androidJar := findAndroidJar()
	if androidJar == "" {
		return fmt.Errorf("android.jar tidak ditemukan di sistem")
	}

	genDir := filepath.Join(cfg.BuildDir, "gen")
	_ = os.MkdirAll(genDir, 0755)

	cmd := exec.Command("aapt", "package",
		"-f",
		"-m",
		"-J", genDir,
		"-M", filepath.Join(cfg.BuildDir, "AndroidManifest.xml"),
		"-S", filepath.Join(cfg.BuildDir, "res"),
		"-I", androidJar,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func compileJava(cfg Config) error {
	androidJar := "/data/data/com.termux/files/usr/share/java/android-24.jar"
	if !fileExists(androidJar) {
		androidJar = findAndroidJar()
	}
	binDir := filepath.Join(cfg.BuildDir, "bin")
	_ = os.MkdirAll(binDir, 0755)

	var javaFiles []string
	_ = filepath.Walk(cfg.BuildDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") {
			javaFiles = append(javaFiles, path)
		}
		return nil
	})

	args := append([]string{
		"-bootclasspath", androidJar,
		"-d", binDir,
	}, javaFiles...)

	cmd := exec.Command("ecj", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runDX(cfg Config) error {
	binDir := filepath.Join(cfg.BuildDir, "bin")
	dexFile := filepath.Join(cfg.BuildDir, "classes.dex")

	cmd := exec.Command("dx", "--dex", "--output="+dexFile, binDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func buildUnsignedAPK(cfg Config, outAPKPath string) error {
	androidJar := findAndroidJar()

	cmd := exec.Command("aapt", "package",
		"-f",
		"-M", filepath.Join(cfg.BuildDir, "AndroidManifest.xml"),
		"-S", filepath.Join(cfg.BuildDir, "res"),
		"-I", androidJar,
		"-F", outAPKPath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("aapt pack apk error: %v", err)
	}

	cmdAdd := exec.Command("aapt", "add", outAPKPath, "classes.dex")
	cmdAdd.Dir = cfg.BuildDir
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	return cmdAdd.Run()
}

func signAPK(cfg Config, inputAPK, outputAPK string) error {
	keystorePath := filepath.Join(cfg.BuildDir, "debug.keystore")

	genKey := exec.Command("keytool", "-genkeypair",
		"-keystore", keystorePath,
		"-alias", "androiddebugkey",
		"-storepass", "android",
		"-keypass", "android",
		"-keyalg", "RSA",
		"-keysize", "2048",
		"-validity", "10000",
		"-dname", "CN=Android Debug,O=Android,C=US",
	)
	_ = genKey.Run()

	cmdSign := exec.Command("apksigner", "sign",
		"--ks", keystorePath,
		"--ks-pass", "pass:android",
		"--key-pass", "pass:android",
		"--ks-key-alias", "androiddebugkey",
		"--out", outputAPK,
		inputAPK,
	)
	cmdSign.Stdout = os.Stdout
	cmdSign.Stderr = os.Stderr
	return cmdSign.Run()
}

func findAndroidJar() string {
	paths := []string{
		"/system/framework/framework-res.apk",
		"/data/data/com.termux/files/usr/share/java/android-24.jar",
		"/data/data/com.termux/files/usr/share/java/android.jar",
		"/usr/share/java/android.jar",
	}
	for _, p := range paths {
		if fileExists(p) {
			return p
		}
	}
	return ""
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
