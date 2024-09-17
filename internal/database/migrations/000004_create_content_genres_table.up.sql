CREATE TABLE IF NOT EXISTS `content_genres` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `content_id` int(11) NOT NULL,
    `genre_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`content_id`) REFERENCES `contents`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`genre_id`) REFERENCES `genres`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;