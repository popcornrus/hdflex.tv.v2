CREATE TABLE IF NOT EXISTS `content_countries` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `content_id` int(11) NOT NULL,
    `country_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;