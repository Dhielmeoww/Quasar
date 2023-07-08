CREATE TABLE `article` (
  `id` integer PRIMARY KEY,
  `title` varchar(255),
  `body` text,
  `status` varchar(255),
  `category_id` integer,
  `publisher_id` integer
);

CREATE TABLE `kategori` (
  `id` integer PRIMARY KEY,
  `category_name` varchar(20)
);

CREATE TABLE `publisher` (
  `id` integer,
  `publisher_name` varchar(255)
);

ALTER TABLE `article` ADD FOREIGN KEY (`category_id`) REFERENCES `kategori` (`id`);

ALTER TABLE `article` ADD FOREIGN KEY (`publisher_id`) REFERENCES `publisher` (`id`);
