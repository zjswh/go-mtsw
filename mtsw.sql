CREATE TABLE `mtsw_mail` (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `title` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商城名称',
                             `aid` int(11) NOT NULL DEFAULT '0',
                             `uin` int(11) NOT NULL,
                             `intro` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商城简介',
                             `shareImg` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分享图片',
                             `contactNumber` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系电话',
                             `contactNumberDefault` tinyint(1) NOT NULL COMMENT '联系电话是否设为默认',
                             `qrCode` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客户二维码',
                             `qrCodeDefault` tinyint(1) NOT NULL COMMENT '二维码是否设为默认',
                             `createTime` int(11) NOT NULL DEFAULT '0',
                             `updateTime` int(11) NOT NULL DEFAULT '0',
                             PRIMARY KEY (`id`) USING BTREE,
                             KEY `i_uin` (`uin`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `mtsw_shipping_template` (
                                          `id` int(11) NOT NULL AUTO_INCREMENT,
                                          `uin` int(11) NOT NULL,
                                          `aid` int(11) NOT NULL DEFAULT '0',
                                          `title` varchar(10) NOT NULL,
                                          `type` tinyint(1) NOT NULL COMMENT '1包邮 0不包邮',
                                          `baseConfig` text COMMENT '配置',
                                          `extraConfig` text,
                                          `limitArea` text COMMENT '限制购买地区',
                                          `deleted` tinyint(1) NOT NULL DEFAULT '0',
                                          `createTime` datetime DEFAULT NULL,
                                          `updateTime` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
                                          PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

CREATE TABLE `goods` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `uin` int(11) DEFAULT NULL,
                         `aid` int(11) DEFAULT NULL,
                         `name` varchar(100) NOT NULL,
                         `desc` varchar(500) DEFAULT NULL COMMENT '商品描述',
                         `price` decimal(6,2) DEFAULT NULL,
                         `shop_price` decimal(11,2) DEFAULT NULL COMMENT '划线价格',
                         `stock` int(11) DEFAULT '0' COMMENT '库存',
                         `sum` int(11) DEFAULT '0' COMMENT '总数',
                         `source` int(4) DEFAULT '1' COMMENT '商品来源 1链接2口令',
                         `sourceContent` text COMMENT '来源内容',
                         `context` text,
                         `limitNum` int(11) DEFAULT '0' COMMENT '限购数',
                         `orderNum` int(11) DEFAULT NULL COMMENT '排序',
                         `type` tinyint(1) DEFAULT '1' COMMENT '1 =>基础类型2外链',
                         `freight` decimal(10,2) DEFAULT '0.00' COMMENT '运费',
                         `freightTemplateId` int(11) DEFAULT '0' COMMENT '运费模板',
                         `freightType` tinyint(255) DEFAULT '3' COMMENT '1直付2模板3包邮',
                         `pickUpAddress` varchar(60) DEFAULT '' COMMENT '自提地址',
                         `pickUpDeadline` int(4) DEFAULT '1' COMMENT '自提期限',
                         `contactNumber` varchar(20) DEFAULT NULL COMMENT '联系电话',
                         `qrCode` varchar(150) DEFAULT NULL COMMENT '客服二维码',
                         `isTop` tinyint(1) DEFAULT '0' COMMENT '1->推荐商品',
                         `status` int(11) DEFAULT '0' COMMENT '1->正常上架',
                         `createTime` int(11) DEFAULT NULL,
                         `updateTime` int(11) DEFAULT NULL,
                         `deleteTime` int(11) DEFAULT NULL,
                         `categoryId` int(11) NOT NULL DEFAULT '0' COMMENT '分类id',
                         `launchTime` int(11) NOT NULL DEFAULT '0' COMMENT '上架时间',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=565 DEFAULT CHARSET=utf8;

CREATE TABLE `goods_bind_tag` (
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `goodsId` int(11) NOT NULL COMMENT '商品id',
                                  `tagId` int(11) NOT NULL COMMENT '标签id',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `goods_imgs` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `goodsId` int(11) DEFAULT NULL,
                              `img` varchar(250) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=517 DEFAULT CHARSET=utf8;

CREATE TABLE `goods_product_specs` (
                                       `id` int(11) NOT NULL AUTO_INCREMENT,
                                       `goodsId` int(11) NOT NULL,
                                       `specs` varchar(150) DEFAULT NULL COMMENT '规格',
                                       `stock` int(11) NOT NULL COMMENT '数量',
                                       `sum` int(11) NOT NULL DEFAULT '0' COMMENT '总数',
                                       `price` decimal(11,2) DEFAULT NULL COMMENT '价格',
                                       `img` varchar(200) DEFAULT '' COMMENT '图片',
                                       `createTime` int(10) NOT NULL,
                                       `updateTime` int(10) DEFAULT NULL,
                                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1808 DEFAULT CHARSET=utf8;

CREATE TABLE `goods_show_case` (
                                   `id` int(11) NOT NULL AUTO_INCREMENT,
                                   `uin` int(11) NOT NULL,
                                   `title` varchar(10) NOT NULL COMMENT '橱窗标题',
                                   `aid` int(11) NOT NULL DEFAULT '0' COMMENT '频道标识',
                                   `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1正常0删除',
                                   `createTime` int(11) NOT NULL,
                                   `updateTime` int(11) NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8;

CREATE TABLE `goods_spec` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `goodsId` int(11) NOT NULL COMMENT '商品id',
                              `name` varchar(100) NOT NULL,
                              `deleteTime` int(11) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=542 DEFAULT CHARSET=utf8;

CREATE TABLE `goods_spec_value` (
                                    `id` int(11) NOT NULL AUTO_INCREMENT,
                                    `goodsId` int(11) DEFAULT NULL,
                                    `specId` int(11) NOT NULL,
                                    `name` varchar(500) NOT NULL,
                                    `deleteTime` int(11) DEFAULT NULL,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1508 DEFAULT CHARSET=utf8;

CREATE TABLE `goods_tag` (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
                             `alias` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT '',
                             `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '0为系统预设 1为客户自定义',
                             `createTime` int(11) NOT NULL,
                             `updateTime` int(11) DEFAULT '0',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `menu_goods` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `include_id` int(11) DEFAULT '0',
                              `menu_id` int(11) NOT NULL,
                              `goods_id` int(11) NOT NULL,
                              `sort` int(5) NOT NULL DEFAULT '99' COMMENT '排序',
                              `type` int(3) NOT NULL DEFAULT '1' COMMENT '1为菜单下 2为购物袋',
                              `isTop` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                              `oldSort` int(5) DEFAULT '0',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1258 DEFAULT CHARSET=utf8;

CREATE TABLE `mtsw_category` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `uin` int(11) NOT NULL,
                                 `aid` int(11) NOT NULL DEFAULT '0',
                                 `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
                                 `sort` int(7) NOT NULL DEFAULT '99' COMMENT '排序',
                                 `createTime` int(11) NOT NULL DEFAULT '0',
                                 `updateTime` int(11) NOT NULL DEFAULT '0',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 KEY `i_uin` (`uin`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `orders` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `uin` int(11) NOT NULL,
                          `aid` int(11) DEFAULT '0',
                          `transactionId` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                          `orderNo` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                          `trackingCompany` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '物流公司',
                          `trackingNo` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT '',
                          `refundId` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '退款单号',
                          `refundReason` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '退款原因',
                          `refundType` int(5) NOT NULL DEFAULT '0' COMMENT '退款类型',
                          `userId` int(11) DEFAULT NULL,
                          `status` int(11) NOT NULL DEFAULT '0' COMMENT '0->未支付 10->已支付 11->支付成功但是订单失败',
                          `num` int(11) NOT NULL DEFAULT '0',
                          `price` decimal(11,2) DEFAULT NULL,
                          `payType` int(11) NOT NULL DEFAULT '0' COMMENT '支付方式 1=>微信支付',
                          `code` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '核销码',
                          `createTime` int(11) DEFAULT NULL,
                          `updateTime` int(11) DEFAULT NULL,
                          `deleteTime` int(11) DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `userId` (`userId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=493 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `orders_detail` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `goodsId` int(11) DEFAULT NULL,
                                 `goodsImg` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `goodsName` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `goodsPrice` decimal(11,2) DEFAULT NULL,
                                 `freight` decimal(10,2) DEFAULT '0.00' COMMENT '运费',
                                 `orderId` int(11) DEFAULT NULL,
                                 `orderMsg` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '买家留言',
                                 `specIds` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `newSpecId` int(11) DEFAULT NULL,
                                 `specValues` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `name` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `mobile` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `province` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `city` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `area` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `detail` varchar(400) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `freightType` int(4) DEFAULT '1' COMMENT '配送方式',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=493 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `orders_refund_record` (
                                        `id` int(11) NOT NULL AUTO_INCREMENT,
                                        `uin` int(11) NOT NULL,
                                        `orderId` int(11) NOT NULL,
                                        `reason` varchar(150) NOT NULL DEFAULT '' COMMENT '退款原因',
                                        `type` int(5) NOT NULL COMMENT '退款类型',
                                        `createTime` datetime NOT NULL,
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

CREATE TABLE `showcase_bind_goods` (
                                       `id` int(11) NOT NULL AUTO_INCREMENT,
                                       `caseId` int(11) NOT NULL COMMENT '橱窗id',
                                       `goodsId` int(11) NOT NULL COMMENT '商品id',
                                       `deleted` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否移除 1移除',
                                       `isTop` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否置顶 1置顶',
                                       `isExplaining` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否讲解中 1是',
                                       `sort` int(10) NOT NULL COMMENT '排序',
                                       `createTime` int(11) NOT NULL,
                                       `updateTime` int(11) NOT NULL DEFAULT '0',
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8;

CREATE TABLE `showcase_include` (
                                    `id` int(11) NOT NULL AUTO_INCREMENT,
                                    `include` varchar(30) NOT NULL COMMENT '引用id',
                                    `includeName` varchar(180) NOT NULL COMMENT '引用名称',
                                    `caseId` int(11) NOT NULL COMMENT '橱窗id',
                                    `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否开启 1开0关',
                                    `createTime` int(11) NOT NULL,
                                    `updateTime` int(11) NOT NULL DEFAULT '0',
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;
