// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

// #include <stdlib.h>
// #include <string.h>
// #include <libudev.h>
// #include "utils_udev.h"

// char *new_str(const char *src_str) {
//         char *dest_str = NULL;
//         if (src_str != NULL) {
//                 int n = strlen(src_str);
//                 dest_str = (char*)malloc(n+1);
//                 memcpy(dest_str, src_str, n);
//                 dest_str[n] = '\0';
//         }
//         return dest_str;
// }

// char *get_device_product(const char *syspath) {
//         struct udev *udev;
//         struct udev_device *device;
//         const char *product_tmp;
//         char *vendor = NULL;

//         udev = udev_new();
//         if (udev == NULL) {
//                 return vendor;
//         }

//         device = udev_device_new_from_syspath(udev, syspath);
//         if (device == NULL) {
//                 udev_unref(udev);
//                 return vendor;
//         }

//         product_tmp = udev_device_get_property_value(device, "ID_MODEL_FROM_DATABASE");
//         if (!product_tmp) {
//                 /* Sometimes ID_PRODUCT_FROM_DATABASE is used? */
//                 product_tmp = udev_device_get_property_value(device, "ID_PRODUCT_FROM_DATABASE");
//         }
//         vendor = new_str(product_tmp);

//         udev_device_unref(device);
//         udev_unref(udev);
//         return vendor;
// }

// char *get_device_vendor(const char *syspath) {
//         struct udev *udev;
//         struct udev_device *device;
//         const char *vendor_tmp;
//         char *vendor = NULL;

//         udev = udev_new();
//         if (udev == NULL) {
//                 return vendor;
//         }

//         device = udev_device_new_from_syspath(udev, syspath);
//         if (device == NULL) {
//                 udev_unref(udev);
//                 return vendor;
//         }

//         vendor_tmp = udev_device_get_property_value(device, "ID_VENDOR_FROM_DATABASE");
//         vendor = new_str(vendor_tmp);

//         udev_device_unref(device);
//         udev_unref(udev);
//         return vendor;
// }

// int is_device_has_property(struct udev_device *device, const char *property) {
//         int ret = -1;
//         struct udev_list_entry *list_entry;
//         const char *owned_prop;
//         udev_list_entry_foreach(list_entry, udev_device_get_properties_list_entry(device)) {
//                 owned_prop = udev_list_entry_get_name(list_entry);
//                 if (strncmp(property, owned_prop, strlen(owned_prop)) == 0) {
//                         ret = 0;
//                         break;
//                 }
//         }
//         return ret;
// }

// int is_usb_device(const char *syspath) {
//         struct udev *udev;
//         struct udev_device *device;
//         const char *id_bus;
//         int ret = -1;

//         udev = udev_new();
//         if (udev == NULL) {
//                 return ret;
//         }

//         device = udev_device_new_from_syspath(udev, syspath);
//         if (device == NULL) {
//                 udev_unref(udev);
//                 return ret;
//         }

//         id_bus = udev_device_get_property_value(device, "ID_BUS");
//         if (id_bus != NULL && strncmp(id_bus, "usb", strlen(id_bus)) == 0) {
//                 ret = 0;
//         }

//         udev_device_unref(device);
//         udev_unref(udev);
//         return ret;
// }
