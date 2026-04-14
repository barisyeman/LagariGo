# LagariGo

> Hız, sadelik, güven. Laravel ergonomisi + Go performansı.

PHP/Laravel kolaylığını Go'nun tip güvenliği ve hızıyla birleştiren bir başlangıç şablonu. Klonla, kendi temanı giy, devam et.

## Stack

- **Fiber** — HTTP framework
- **Templ** — tip güvenli HTML şablonları (derleme zamanı kontrolü)
- **HTMX** — sunucu taraflı reaktivite, JS framework yok
- **GORM** — SQLite veya MySQL
- **Düz CSS** — kendi temanı getir, Tailwind yok

## Kurulum

```bash
# 1. Templ CLI (bir kere)
make install-tools

# 2. Bağımlılıklar
go mod tidy

# 3. Konfigürasyon
cp .env.example .env

# 4. Çalıştır
make dev
```

`http://localhost:3000` adresinde açılır.

İlk açılışta otomatik olarak:
- SQLite veritabanı oluşturulur (`./lagarigo.db`)
- Admin kullanıcı: `.env` içindeki `ADMIN_EMAIL` / `ADMIN_PASSWORD` (varsayılan `admin@lagarigo.local` / `admin123`)
- Örnek menü linkleri ve bir hoş geldin sayfası

## Klasör Yapısı

```
cmd/server/         # main.go (giriş noktası)
internal/
  auth/             # session yönetimi
  config/           # .env loader
  database/         # GORM modelleri + bağlantı + seed
  handler/          # HTTP handler'lar (controller)
  middleware/       # auth/admin guard
public/assets/      # statik (CSS/JS/img) — kendi temanı buraya
views/
  layouts/          # base.templ
  components/       # header, footer
  pages/            # home, about, contact, login, register, dynamic, 404
  pages/admin/      # admin paneli sayfaları
.env.example
Makefile            # make dev | build | gen
```

## Sayfalar

- `/` — Ana sayfa
- `/about-us` — Hakkımızda
- `/contact` — İletişim
- `/login`, `/register`, `/logout`
- `/admin` — Yönetim paneli (sadece admin)
- `/:slug` — Dinamik sayfalar (admin panelden sınırsız oluşturulabilir)

> Reserved slug'lar (`about-us`, `contact`, `login`, `register`, `admin`, `assets`) dinamik sayfa olarak kullanılamaz.

## Veritabanı

`.env` içinde `DB_DRIVER=sqlite` (varsayılan) veya `DB_DRIVER=mysql` ile değiştir. MySQL için `DB_HOST`, `DB_USER`, vb. doldur.

## Güvenlik

- **XSS**: Templ tüm değişkenleri otomatik escape eder. `templ.SafeURL` sadece güvenli URL'ler için kullanılır.
- **CSRF**: Tüm POST formlarında `_csrf` token bulunur, Fiber CSRF middleware ile doğrulanır.
- **Şifre**: bcrypt ile hash'lenir.
- **Session**: HttpOnly + SameSite=Lax cookie.

## Tema Değiştirme

`public/assets/css/style.css` dosyasını sil veya değiştir. CSS sınıfları semantiktir (`.btn`, `.card`, `.flash`, vb.) — kendi tasarım sistemine kolayca adapte edilir.

## Komutlar

```bash
make dev        # Templ generate + go run
make build      # Templ generate + go build → bin/lagarigo
make gen        # Sadece templ generate
make tidy       # go mod tidy
make clean      # bin/ + *.db + _templ.go dosyalarını sil
```

## Geliştirme

Templ dosyalarını (`.templ`) düzenledikten sonra `make gen` veya `make dev` ile yeniden üretmen gerekir. Veya `templ generate --watch` ile otomatik üretim yapabilirsin.

## Lisans

MIT
