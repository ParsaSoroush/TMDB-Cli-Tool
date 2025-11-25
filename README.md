# TMDB CLI – Movie Fetcher

A simple Go command-line application to fetch movies from The Movie Database (TMDB) API and show them in the terminal.

[Project Page](https://roadmap.sh/projects/tmdb-cli)

---

## 🔧 Requirements
- Go 1.20+
- A TMDB API Key → https://www.themoviedb.org/

---

## 📌 Installation

```bash
git clone <your-repo-url>
cd tmdb-cli
go mod init tmdb-app
go build -o tmdb-app
```

## 🔑 Set API Key
```bash
export TMDB_API_KEY="YOUR_API_KEY"
```

## ▶ Usage
```bash
./tmdb-app --type "playing"   # Now Playing
./tmdb-app --type "popular"   # Popular Movies
./tmdb-app --type "top"       # Top Rated Movies
./tmdb-app --type "upcoming"  # Upcoming Movies
```