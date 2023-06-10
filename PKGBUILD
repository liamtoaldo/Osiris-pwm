# This is an example PKGBUILD file. Use this as a start to creating your own,
# and remove these comments. For more information, see 'man PKGBUILD'.
# NOTE: Please fill out the license field for your package! If it is unknown,
# then please put 'unknown'.

# Maintainer: liamtoaldo <mailytdlg7@gmail.com>
pkgname="Osiris-pwm"
pkgver=1.1
pkgrel=1
pkgdesc="A simple and lightweight encrypted password manager written in Go."
arch=("x86_64")
license=('MIT')
source=("https://github.com/liamtoaldo/Osiris-pwm.git")

build() {
    if [[ -d "Osiris-pwm" ]]
    then
        echo "Directory already exists, removing the previous version"
        rm -rf "Osiris-pwm"
    fi
    git clone "https://github.com/liamtoaldo/Osiris-pwm.git"
	cd "Osiris-pwm/"
	go build
}
package() {
    cd "Osiris-pwm"
    echo "Installation complete"
}
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
md5sums=('f30fbe7b5b77dce2286aa1e96011e1ec')
