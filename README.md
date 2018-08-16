# Bytom-Mobile-Wallet-SDK

## Prepare
   Make sure you have installed Golang/JDK/Android SDK/Android NDK/Xcode correctly.
   - Install Gomobile:
   ```shell
   $ go get golang.org/x/mobile/cmd/gomobile
   $ gomobile init -ndk ~/path/to/your/ndk
   ```
   At least Go 1.7 is required. For detailed instructions, see https://golang.org/wiki/Mobile.

   - Clone this project to your $GOPATH/src
   ```shell
    git clone https://github.com/Bytom-Community/Bytom-Mobile-Wallet-SDK $GOPATH/src/github.com/bytom-community/mobile
   ```

## Build
   ```shell
   cd $GOPATH/src/github.com/bytom-community/mobile
   ```
   - Android
   ```shell
   $ gomobile bind -target=android -ldflags=-s github.com/bytom-community/mobile/sdk/
   ```

   - iOS
   ```shell
   $ gomobile bind -target=ios -ldflags=-w github.com/bytom-community/mobile/sdk/
   ```
