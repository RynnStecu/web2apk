package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("📦 WEB2APK DEPENDENCY & SYSTEM SETUP INSTALLER 📦")
	fmt.Println("==================================================")

	var pkgCmd string
	var packages []string

	// Deteksi Lingkungan Sistem
	if fileExists("/data/data/com.termux/files/usr/bin/pkg") {
		fmt.Println("🐧 Lingkungan Terdeteksi: Termux (Android)")
		pkgCmd = "pkg"
		packages = []string{
			"golang",
			"openjdk-21",
			"ecj",
			"dx",
			"aapt",
			"apksigner",
			"imagemagick",
		}
	} else if runtime.GOOS == "linux" {
		fmt.Println("🐧 Lingkungan Terdeteksi: Linux / Ubuntu / Debian")
		pkgCmd = "apt-get"
		packages = []string{
			"golang",
			"openjdk-21-jdk",
			"ecj",
			"aapt",
			"apksigner",
			"imagemagick",
		}
	} else {
		fmt.Printf("⚠️  Sistem OS '%s' belum didukung otomatis untuk auto-install paket.\n", runtime.GOOS)
		os.Exit(1)
	}

	fmt.Println("🔄 Memperbarui package list dan menginstal dependensi...")
	fmt.Println("📋 Paket yang diinstal:", packages)
	fmt.Println("--------------------------------------------------")

	// 1. Update Package Manager
	var updateCmd *exec.Cmd
	if pkgCmd == "pkg" {
		updateCmd = exec.Command("pkg", "update", "-y")
	} else {
		updateCmd = exec.Command("sudo", "apt-get", "update", "-y")
	}
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	_ = updateCmd.Run()

	// 2. Install Packages
	args := append([]string{"install", "-y"}, packages...)
	var installCmd *exec.Cmd
	if pkgCmd == "pkg" {
		installCmd = exec.Command("pkg", args...)
	} else {
		args = append([]string{"apt-get"}, args...)
		installCmd = exec.Command("sudo", args...)
	}
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		fmt.Printf("\n❌ Gagal menginstal beberapa paket: %v\n", err)
	} else {
		fmt.Println("\n✅ Semua dependensi sistem berhasil diinstal!")
	}

	// 3. Build & Install binary web2apk ke $HOME/bin atau /usr/local/bin
	fmt.Println("--------------------------------------------------")
	fmt.Println("🔨 Membangun binary executable 'web2apk'...")

	buildCmd := exec.Command("go", "build", "-o", "web2apk", "main.go")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		fmt.Printf("❌ Gagal melakukan 'go build': %v\n", err)
		os.Exit(1)
	}

	// Copy binary ke PATH
	homeDir, _ := os.UserHomeDir()
	targetBin := homeDir + "/bin/web2apk"
	_ = os.MkdirAll(homeDir+"/bin", 0755)

	if err := exec.Command("cp", "web2apk", targetBin).Run(); err == nil {
		_ = exec.Command("chmod", "+x", targetBin).Run()
		fmt.Printf("🚀 Berhasil menginstal 'web2apk' ke: %s\n", targetBin)
		fmt.Println("💡 Anda sekarang bisa mengetik perintah 'web2apk' dari mana saja!")
	}

	fmt.Println("==================================================")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
