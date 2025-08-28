include $(TOPDIR)/rules.mk

PKG_NAME:=luci-app-g
PKG_VERSION:=0.1.0
PKG_RELEASE:=1

LUCI_TITLE:=g - Natural Language CLI for OpenWrt (Web UI)
LUCI_PKGARCH:=all
LUCI_DEPENDS:=+luci-base +luci-compat +g
LUCI_MAINTAINER:=aezizhu

include $(TOPDIR)/feeds/luci/luci.mk

# call BuildPackage - OpenWrt build system will pick this up for LuCI
$(eval $(call BuildPackage,$(PKG_NAME)))


