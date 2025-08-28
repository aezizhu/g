include $(TOPDIR)/rules.mk

PKG_NAME:=g
PKG_VERSION:=0.1.0
PKG_RELEASE:=1

PKG_MAINTAINER:=aezizhu

include $(INCLUDE_DIR)/package.mk

define Package/g
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=Natural language CLI for OpenWrt
  DEPENDS:=+libc
endef

define Package/g/description
  Natural language CLI that translates requests into safe, audited commands.
endef

define Build/Compile
        # Binary should be prebuilt via external Go toolchain; copy during install
endef

define Package/g/install
        $(INSTALL_DIR) $(1)/usr/bin
        $(INSTALL_BIN) ../../dist/g-$(ARCH) $(1)/usr/bin/g || true
endef

$(eval $(call BuildPackage,g))


