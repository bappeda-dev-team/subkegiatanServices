CREATE TABLE tb_subkegiatan(
    id VARCHAR(255) PRIMARY KEY,
    kode_subkegiatan VARCHAR(255) NOT NULL,
    nama_subkegiatan VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tb_subkegiatan_timestamp
BEFORE UPDATE ON tb_subkegiatan
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();