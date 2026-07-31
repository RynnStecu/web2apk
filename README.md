<div align="center">

  # 🚀 WEB2APK BUILDER
  ### *Transform Any Website URL into a Native Android APK*

  <p align="center">
    <a href="https://golang.org"><img src="https://img.shields.io/badge/Language-Golang_1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Golang" /></a>
    <a href="https://android.com"><img src="https://img.shields.io/badge/Platform-Android_13%2B_Ready-3DDC84?style=for-the-badge&logo=android&logoColor=white" alt="Android" /></a>
    <a href="https://api.kyzzz.eu.cc"><img src="https://img.shields.io/badge/API-Rynnstecu_Engine-FF4500?style=for-the-badge&logo=fastapi&logoColor=white" alt="API" /></a>
    <a href="https://github.com/RynnStecu/web2apk/releases"><img src="https://img.shields.io/badge/Download-Binary_v1.0.0-blueviolet?style=for-the-badge&logo=github&logoColor=white" alt="Release" /></a>
  </p>

  <p align="center">
    <b>⚡ High Performance • 📱 Android 13+ Support • 🎨 Auto PNG Converter • 🚀 1-Click Installer</b>
  </p>

  ---

</div>

## 🌟 About The Project

**Web2APK Builder** adalah tools otomatis berbasis **Golang** yang sangat cepat untuk mengubah URL Website apapun menjadi **Aplikasi Native Android (APK)** secara langsung di perangkat Anda tanpa perantara server pihak ketiga.

Dilengkapi dengan integrasi Android Toolchain (`aapt`, `ecj`, `dx`, `apksigner`), tools ini menghasilkan file APK yang siap langsung di-install pada perangkat Android modern (**Android 13 & Android 14 Ready**).

---

## ✨ Features Highlight

```
 ┌───────────────────────────────────────────────────────────────────────┐
 │  📱 Android 13+ Ready    │ targetSdkVersion 34, Granular Media Perms │
 │  📁 HTML5 File Upload    │ Integrated WebChromeClient FilePicker      │
 │  🖼️ Auto Image Convert   │ Auto-convert JPG/JPEG icon to Android PNG │
 │  ⚡ Instant 1-Click Run  │ Direct binary execution without extra steps│
 │  🔑 Auto-Signing Engine  │ Auto Debug Keystore Generator & APKSigner  │
 └───────────────────────────────────────────────────────────────────────┘
```

---

## ⚡ 1-Click All-in-One Command (Selalu Jalan)

Cukup salin dan jalankan **1 baris perintah di bawah ini** untuk menginstal semua dependensi, mengunduh biner `web2apk`, dan memasangnya langsung ke sistem Anda:

### 📱 Termux (Android)
```bash
pkg update && pkg install -y golang openjdk-21 ecj dx aapt apksigner imagemagick curl && mkdir -p $HOME/bin && curl -L -o $HOME/bin/web2apk https://github.com/RynnStecu/web2apk/releases/download/v1.0.0/web2apk && chmod +x $HOME/bin/web2apk
```

### 🐧 Linux (Ubuntu / Debian / Kali)
```bash
sudo apt update && sudo apt install -y golang openjdk-21-jdk ecj aapt apksigner imagemagick curl && sudo mkdir -p /usr/local/bin && sudo curl -L -o /usr/local/bin/web2apk https://github.com/RynnStecu/web2apk/releases/download/v1.0.0/web2apk && sudo chmod +x /usr/local/bin/web2apk
```

---

## 🚀 Quick Usage

Setelah menjalankan perintah installer di atas, Anda bisa langsung mengetik `web2apk` dari folder mana saja tanpa memasukkan string key apapun!

### ⚙️ Command Format

```bash
web2apk -url <TARGET_URL> -name "<APP_NAME>" -package "<PACKAGE_NAME>" -icon "<ICON_PATH>" -out "<OUTPUT_DIR>"
```

### 💡 Example Commands

```bash
# Basic Build
web2apk -url https://api.kyzzz.eu.cc -name "Kyu Apis"

# Full Custom Build with JPG/PNG Icon
web2apk -url https://api.kyzzz.eu.cc \
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

---

## 👨‍💻 Creator & Community Contacts

<div align="center">

| Channel | Link / Address |
| :--- | :--- |
| 👤 **Creator** | **Rynnstecu** |
| ✈️ **Telegram** | [@kyuugaprawan](https://t.me/kyuugaprawan) |
| 📢 **WhatsApp Channel** | [Join Community Channel](https://whatsapp.com/channel/0029Vb7gcbuLdQelWzrTzD3D) |
| 🌐 **API Web Service** | [https://api.kyzzz.eu.cc](https://api.kyzzz.eu.cc) |
| 📦 **GitHub Releases** | [v1.0.0 Release](https://github.com/RynnStecu/web2apk/releases/tag/v1.0.0) |

</div>

<br>

<div align="center">

---
**Made with ❤️ by Rynnstecu** • Powered by Golang & Android Open Source Project

</div>
