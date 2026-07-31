<div align="center">

  # 🚀 WEB2APK BUILDER
  ### *Transform Any Website URL into a Native Android APK*

  <p align="center">
    <a href="https://golang.org"><img src="https://img.shields.io/badge/Language-Golang_1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Golang" /></a>
    <a href="https://android.com"><img src="https://img.shields.io/badge/Platform-Android_13%2B_Ready-3DDC84?style=for-the-badge&logo=android&logoColor=white" alt="Android" /></a>
    <a href="https://api.kyzzz.eu.cc"><img src="https://img.shields.io/badge/API-Rynnstecu_Engine-FF4500?style=for-the-badge&logo=fastapi&logoColor=white" alt="API" /></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-Protected-blueviolet?style=for-the-badge" alt="License" /></a>
  </p>

  <p align="center">
    <b>⚡ High Performance • 📱 Android 13+ Support • 🎨 Auto PNG Converter</b>
  </p>

  ---

</div>

## 🌟 About The Project

**Web2APK Builder** adalah tools otomatis berbasis **Golang** yang sangat cepat untuk mengubah URL Website apapun menjadi **Aplikasi Native Android (APK)** secara langsung di perangkat Anda tanpa perantara server pihak ketiga.

Dilengkapi dengan integrasi Android Toolchain (`aapt`, `ecj`, `dx`, `apksigner`) dan proteksi keamanan tingkat tinggi, tools ini menghasilkan file APK yang siap langsung di-install pada perangkat Android modern (**Android 13 & Android 14 Ready**).

---

## ✨ Features Highlight

```
 ┌───────────────────────────────────────────────────────────────────────┐
 │  📱 Android 13+ Ready    │ targetSdkVersion 34, Granular Media Perms │
 │  📁 HTML5 File Upload    │ Integrated WebChromeClient FilePicker      │
 │  🖼️ Auto Image Convert   │ Auto-convert JPG/JPEG icon to Android PNG │
 │  🔒 Anti-Recode Shield   │ SHA-256 Hash Integrity System              │
 │  🔑 Auto-Signing Engine  │ Auto Debug Keystore Generator & APKSigner  │
 └───────────────────────────────────────────────────────────────────────┘
```

---

## 🛠️ Stack & Dependencies Installation

Sebelum menjalankan **Web2APK Builder**, pastikan dependensi berikut sudah terpasang:

### 📱 Termux (Android)
```bash
pkg update && pkg install -y golang openjdk-21 ecj dx aapt apksigner imagemagick
```

### 🐧 Linux (Ubuntu / Debian / Kali)
```bash
sudo apt update && sudo apt install -y golang openjdk-21-jdk ecj aapt apksigner imagemagick
```

---

## 🚀 Quick Usage

### ⚙️ Command Format

```bash
web2apk -url <TARGET_URL> -name "<APP_NAME>" -package "<PACKAGE_NAME>" -icon "<ICON_PATH>" -out "<OUTPUT_DIR>"
```

### 💡 Example Commands

```bash
# Basic Build
./web2apk -url https://api.kyzzz.eu.cc -name "Kyu Apis"

# Full Custom Build with JPG/PNG Icon
./web2apk -url https://api.kyzzz.eu.cc \
          -name "Kyu Apis" \
          -package "com.kyzzz.apis" \
          -icon "test.jpg" \
          -out "/sdcard/"
```

---

## 🎛️ Command Line Options

| Flag | Description | Default Value | Required |
| :--- | :--- | :--- | :---: |
| `-url` | Target Website URL to convert | `""` | **YES** |
| `-name` | Android Application Name | `"MyWebApp"` | NO |
| `-package` | Android Package Identifier | `"com.mycompany.webapp"` | NO |
| `-icon` | Path to Application Icon (`.png` / `.jpg`) | System Default | NO |
| `-out` | Output directory for `.apk` file | `"."` | NO |
| `-key` | Anti-Recode Integrity Key | `"kyu?-sayang-hasna"` | NO |

---

## 👨‍💻 Creator & Community Contacts

<div align="center">

| Channel | Link / Address |
| :--- | :--- |
| 👤 **Creator** | **Rynnstecu** |
| ✈️ **Telegram** | [@kyuugaprawan](https://t.me/kyuugaprawan) |
| 📢 **WhatsApp Channel** | [Join Community Channel](https://whatsapp.com/channel/0029Vb7gcbuLdQelWzrTzD3D) |
| 🌐 **API Web Service** | [https://api.kyzzz.eu.cc](https://api.kyzzz.eu.cc) |

</div>

<br>

<div align="center">

---
**Made with ❤️ by Rynnstecu** • Powered by Golang & Android Open Source Project

</div>
