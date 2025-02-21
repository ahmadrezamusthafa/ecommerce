-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
INSERT INTO products (name, description, price, created_at, updated_at)
VALUES ('Mouse Wireless', 'Mouse wireless ergonomis dengan akurasi tinggi', 250000, NOW(), NOW()),
       ('Keyboard Mekanikal', 'Keyboard mekanikal RGB dengan switch biru', 1200000, NOW(), NOW()),
       ('Monitor Gaming', 'Monitor gaming 27 inci, refresh rate 165Hz, waktu respons 1ms', 4500000, NOW(), NOW()),
       ('Hub USB-C', 'Hub USB-C 6-in-1 dengan HDMI, USB 3.0, dan pembaca kartu SD', 550000, NOW(), NOW()),
       ('Headphone Bluetooth', 'Headphone over-ear dengan fitur noise-canceling', 1750000, NOW(), NOW()),
       ('Jam Pintar', 'Jam pintar dengan fitur pelacak kebugaran dan monitor detak jantung', 2300000, NOW(), NOW()),
       ('SSD Eksternal', 'SSD eksternal 1TB dengan koneksi USB 3.2', 2100000, NOW(), NOW()),
       ('Speaker Portable', 'Speaker Bluetooth waterproof dengan bass yang dalam', 800000, NOW(), NOW()),
       ('Charger Nirkabel', 'Charger nirkabel cepat untuk berbagai merek smartphone', 350000, NOW(), NOW()),
       ('Stand Laptop', 'Stand laptop aluminium dengan ventilasi pendingin', 450000, NOW(), NOW()),
       ('Kamera CCTV', 'Kamera CCTV full HD dengan fitur night vision', 1200000, NOW(), NOW()),
       ('Microphone Kondensor', 'Microphone kondensor untuk streaming dan podcasting', 650000, NOW(), NOW()),
       ('Power Bank', 'Power bank 20000mAh dengan fitur fast charging', 500000, NOW(), NOW()),
       ('Tablet Grafis', 'Tablet grafis dengan stylus pen sensitivitas tinggi', 1800000, NOW(), NOW()),
       ('Tripod Kamera', 'Tripod kamera dengan tinggi 1,5m dan kepala putar 360°', 350000, NOW(), NOW()),
       ('Lampu LED RGB', 'Lampu LED RGB dengan kontrol warna melalui aplikasi', 250000, NOW(), NOW()),
       ('Mousepad Gaming', 'Mousepad gaming dengan permukaan mikro-texture', 150000, NOW(), NOW()),
       ('Keyboard Nirkabel', 'Keyboard nirkabel dengan desain minimalis', 450000, NOW(), NOW()),
       ('Smart TV 43"', 'Smart TV 43 inci dengan resolusi 4K dan fitur Android TV', 4500000, NOW(), NOW()),
       ('Kipas Pendingin Laptop', 'Cooling pad dengan 5 kipas dan LED biru', 200000, NOW(), NOW()),
       ('Flashdisk 128GB', 'Flashdisk USB 3.1 dengan kapasitas 128GB', 200000, NOW(), NOW()),
       ('Tas Laptop', 'Tas laptop tahan air dengan banyak kompartemen', 350000, NOW(), NOW()),
       ('Mouse Ergonomis', 'Mouse ergonomis untuk mengurangi ketegangan pergelangan tangan', 300000, NOW(), NOW()),
       ('Ring Light', 'Ring light LED 12 inci dengan tripod dan pengaturan kecerahan', 400000, NOW(), NOW()),
       ('Hard Disk Eksternal', 'Hard disk eksternal 2TB dengan kecepatan tinggi', 1400000, NOW(), NOW()),
       ('Game Controller', 'Game controller nirkabel kompatibel dengan PC dan konsol', 750000, NOW(), NOW()),
       ('Kabel HDMI 4K', 'Kabel HDMI 4K ultra high-speed sepanjang 2 meter', 120000, NOW(), NOW()),
       ('Headset Gaming', 'Headset gaming dengan surround sound 7.1', 950000, NOW(), NOW()),
       ('Mini Projector', 'Mini projector portabel dengan resolusi Full HD', 2300000, NOW(), NOW()),
       ('Printer Inkjet', 'Printer inkjet multifungsi dengan fitur wireless', 1500000, NOW(), NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
