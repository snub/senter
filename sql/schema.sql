CREATE TABLE sensor_controller (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    mac_address CHAR(12) NOT NULL,
    last_startup DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    CONSTRAINT pk_sensor_controller_id PRIMARY KEY (id),
    CONSTRAINT un_mac_address UNIQUE INDEX (mac_address)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;

CREATE TABLE sensor_controller_config (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    controller_id INT UNSIGNED NOT NULL,
    ip_address VARCHAR(45) NOT NULL,
    update_interval INT UNSIGNED NOT NULL,
    ntp_ip_address VARCHAR(45) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    CONSTRAINT pk_sensor_controller_config_id PRIMARY KEY (id),
    CONSTRAINT un_controller_id UNIQUE INDEX (controller_id),
    CONSTRAINT fk_controller_config_id FOREIGN KEY (controller_id) REFERENCES sensor_controller (id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;


CREATE TABLE sensor (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    device_address CHAR(16) NOT NULL,
    name VARCHAR(64) NULL DEFAULT NULL,
    description VARCHAR(255) NULL DEFAULT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    CONSTRAINT pk_sensor_id PRIMARY KEY (id),
    CONSTRAINT un_device_address UNIQUE INDEX (device_address)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;

CREATE TABLE sensor_temperature (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    sensor_id INT UNSIGNED NOT NULL,
    timestamp DATETIME NOT NULL,
    value DECIMAL(5, 2) NOT NULL,
    CONSTRAINT pk_sensor_temperature_id PRIMARY KEY (id),
    INDEX id_sensor_id (sensor_id),
    INDEX id_timestamp (timestamp),
    CONSTRAINT fk_sensor_id FOREIGN KEY (sensor_id) REFERENCES sensor (id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;
