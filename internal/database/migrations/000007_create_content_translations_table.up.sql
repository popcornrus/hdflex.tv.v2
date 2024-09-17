CREATE TABLE IF NOT EXISTS `content_translations` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `content_id` int(11) NOT NULL,
    `translation_id` int(11) NOT NULL,
    `quality` varchar(128) NOT NULL,
    `max_quality` int(4) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`content_id`) REFERENCES `contents`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`translation_id`) REFERENCES `translations` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;