package services

import (
    "fmt"
    "time"

    "gorm.io/gorm"
    "evermos/configs"
    "evermos/models"
)

type TrxService struct {
    DB *gorm.DB
}

func NewTrxService() *TrxService {
    return &TrxService{DB: configs.DB}
}

func (s *TrxService) CreateTransaction(userID uint, alamatID uint, items []map[string]int) (*models.Trx, error) {
    var trx *models.Trx
    err := s.DB.Transaction(func(tx *gorm.DB) error {
        total := 0
        var details []models.DetailTrx
        for _, it := range items {
            pid := uint(it["produk_id"])
            qty := it["kuantitas"]
            var prod models.Produk
            if err := tx.Preload("Toko").First(&prod, pid).Error; err != nil {
                return err
            }
            logp := models.LogProduk{
                ProdukID: prod.ID,
                NamaProduk: prod.NamaProduk,
                Slug: prod.Slug,
                HargaReseller: prod.HargaReseller,
                HargaKonsumen: prod.HargaKonsumen,
                Deskripsi: prod.Deskripsi,
                TokoID: prod.TokoID,
                CategoryID: prod.CategoryID,
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            }
            if err := tx.Create(&logp).Error; err != nil {
                return err
            }
            harga := prod.HargaKonsumen * qty
            total += harga
            if prod.Stok < qty {
                return fmt.Errorf("not enough stock for product %d", prod.ID)
            }
            prod.Stok -= qty
            if err := tx.Save(&prod).Error; err != nil {
                return err
            }
            details = append(details, models.DetailTrx{
                LogProdukID: logp.ID,
                TokoID: prod.TokoID,
                Kuantitas: qty,
                HargaTotal: harga,
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            })
        }
        trx = &models.Trx{
            UserID: userID,
            AlamatPengirimanID: alamatID,
            HargaTotal: total,
            KodeInvoice: fmt.Sprintf("INV-%d", time.Now().Unix()),
            MethodBayar: "COD",
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        }
        if err := tx.Create(trx).Error; err != nil {
            return err
        }
        for i := range details {
            details[i].TrxID = trx.ID
            if err := tx.Create(&details[i]).Error; err != nil {
                return err
            }
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    s.DB.Preload("Details").First(trx, trx.ID)
    return trx, nil
}
