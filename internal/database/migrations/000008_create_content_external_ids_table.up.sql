CREATE TABLE IF NOT EXISTS `content_external_ids` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `content_id` int(11) NOT NULL,
    `external_id` varchar(36) NOT NULL,
    `external_type` tinyint(4) NOT NULL,
    `external_rating` float(8, 2) NOT NULL,
    `external_rating_votes` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;