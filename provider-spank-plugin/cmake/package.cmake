if(CMAKE_SIZEOF_VOID_P EQUAL 8)
  set(CPACK_PACKAGE_ARCHITECTURE "x86_64")
else()
  set(CPACK_PACKAGE_ARCHITECTURE "x86")
endif()
set(CPACK_PACKAGE_DESCRIPTION "SPANK plugin for cluster providers")
set(CPACK_PACKAGE_VERSION_MAJOR 1)
set(CPACK_PACKAGE_VERSION_MINOR 0)
set(CPACK_PACKAGE_VERSION_PATCH 5)
set(CPACK_PACKAGE_VENDOR "DeepSquare SA")
set(CPACK_PACKAGING_INSTALL_PREFIX "${CMAKE_INSTALL_PREFIX}")
set(CPACK_PACKAGE_CONTACT "nguyen_marc@live.fr")
set(CPACK_PACKAGE_RELOCATABLE TRUE)
set(CPACK_PACKAGE_FILE_NAME
    "${CMAKE_PROJECT_NAME}-${CPACK_PACKAGE_VERSION_MAJOR}.${CPACK_PACKAGE_VERSION_MINOR}.${CPACK_PACKAGE_VERSION_PATCH}-${CPACK_PACKAGE_ARCHITECTURE}"
)

if(EXISTS "/etc/debian_version")
  set(CPACK_GENERATOR "DEB")
elseif(EXISTS "/etc/redhat-release")
  set(CPACK_GENERATOR "RPM")
else()
  set(CPACK_GENERATOR "TXZ")
endif()

set(CPACK_DEBIAN_PACKAGE_ARCHITECTURE "amd64")
set(CPACK_DEBIAN_PACKAGE_HOMEPAGE "https://github.com/deepsquare-io/the-grid")
set(CPACK_DEBIAN_PACKAGE_SHLIBDEPS OFF)
set(CPACK_DEBIAN_FILE_NAME DEB-DEFAULT)

set(CPACK_RPM_PACKAGE_LICENSE "MIT")
set(CPACK_RPM_PACKAGE_GROUP "Development/Libraries/C")
set(CPACK_RPM_PACKAGE_RELEASE "1")
set(CPACK_RPM_PACKAGE_RELEASE_DIST ON)
set(CPACK_RPM_PACKAGE_URL "https://github.com/deepsquare-io/the-grid")
set(CPACK_RPM_PACKAGE_ARCHITECTURE "x86_64")
set(CPACK_RPM_PACKAGE_AUTOREQ OFF)
set(CPACK_RPM_FILE_NAME RPM-DEFAULT)

include(CPack)
