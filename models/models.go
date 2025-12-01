package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Nama      string    `gorm:"size:255" json:"nama"`
    Email     string    `gorm:"size:255;uniqueIndex" json:"email"`
    Password  string    `gorm:"size:255" json:"-"`
    NoTelp    string    `gorm:"size:255;uniqueIndex" json:"no_telp"`
    IsAdmin   bool      `gorm:"default:false" json:"is_admin"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Toko      Toko      `gorm:"foreignKey:UserID" json:"toko"`
    Alamat    []Alamat  `gorm:"foreignKey:UserID" json:"alamat"`
}

type Toko struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"user_id"`
    NamaToko  string    `gorm:"size:255" json:"nama_toko"`
    UrlFoto   string    `gorm:"size:255" json:"url_foto"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Produk    []Produk  `gorm:"foreignKey:TokoID" json:"produk"`
}

type Alamat struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    UserID         uint      `json:"user_id"`
    JudulAlamat    string    `gorm:"size:255" json:"judul_alamat"`
    NamaPenerima   string    `gorm:"size:255" json:"nama_penerima"`
    NoTelp         string    `gorm:"size:255" json:"no_telp"`
    DetailAlamat   string    `gorm:"type:text" json:"detail_alamat"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

type Category struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Nama      string    `gorm:"size:255" json:"nama"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Produk struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    NamaProduk      string    `gorm:"size:255" json:"nama_produk"`
    Slug            string    `gorm:"size:255" json:"slug"`
    HargaReseller   int       `json:"harga_reseller"`
    HargaKonsumen   int       `json:"harga_konsumen"`
    Stok            int       `json:"stok"`
    Deskripsi       string    `gorm:"type:text" json:"deskripsi"`
    TokoID          uint      `json:"toko_id"`
    CategoryID      uint      `json:"category_id"`
    FotoProduk      []Foto    `gorm:"foreignKey:ProdukID" json:"foto_produk"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type Foto struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    ProdukID  uint      `json:"produk_id"`
    Url       string    `gorm:"size:255" json:"url"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type LogProduk struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    ProdukID      uint      `json:"produk_id"`
    NamaProduk    string    `gorm:"size:255" json:"nama_produk"`
    Slug          string    `gorm:"size:255" json:"slug"`
    HargaReseller int       `json:"harga_reseller"`
    HargaKonsumen int       `json:"harga_konsumen"`
    Deskripsi     string    `gorm:"type:text" json:"deskripsi"`
    TokoID        uint      `json:"toko_id"`
    CategoryID    uint      `json:"category_id"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}

type Trx struct {
    ID               uint        `gorm:"primaryKey" json:"id"`
    UserID           uint        `json:"user_id"`
    AlamatPengirimanID uint      `json:"alamat_pengiriman_id"`
    HargaTotal       int         `json:"harga_total"`
    KodeInvoice      string      `gorm:"size:255" json:"kode_invoice"`
    MethodBayar      string      `gorm:"size:255" json:"method_bayar"`
    Details          []DetailTrx `gorm:"foreignKey:TrxID" json:"details"`
    CreatedAt        time.Time   `json:"created_at"`
    UpdatedAt        time.Time   `json:"updated_at"`
}

type DetailTrx struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    TrxID     uint      `json:"trx_id"`
    LogProdukID uint    `json:"log_produk_id"`
    TokoID    uint      `json:"toko_id"`
    Kuantitas int       `json:"kuantitas"`
    HargaTotal int      `json:"harga_total"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
