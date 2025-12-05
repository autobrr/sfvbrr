# TESTS RUNSHEET

## 1 App

* 1.1 :white_check_mark: A folder containing `setup.zip`, `release.nfo`, and `file_id.diz`.
* 1.2 :white_check_mark: A folder containing `part1.zip`, `part2.zip`, `release.nfo`, and `file_id.diz` (Multiple zips allowed).
* 1.3 :x: A folder missing `file_id.diz` (Violates `min: 1`).
* 1.4 :x: A folder containing two `.nfo` files (Violates `max: 1`).
* 1.5 :x: A folder with `installer.rar` instead of `.zip` (Violates pattern).

## 2 Audiobook

* 2.1 :white_check_mark: A folder with `playlist.m3u`, `checksum.sfv`, `book.nfo`, and 10 `.mp3` files.
* 2.2 :white_check_mark: A folder with multiple `.m3u` playlists (e.g., CD1, CD2) and multiple `.mp3` files.
* 2.3 :x: A folder missing the `.sfv` file (Violates `min: 1`).
* 2.4 :x: A folder containing `.m4b` audio files but no `.mp3` files (Violates `.mp3` `min: 1`).
* 2.5 :x: A folder with two `.nfo` files (Violates `max: 1`).

## 3 Book

* 3.1 :white_check_mark: A folder containing `novel.zip`, `info.nfo`, and `file_id.diz`.
* 3.2 :x: A folder missing the `.nfo` file (Violates `min: 1`).
* 3.3 :x: A folder containing two `file_id.diz` files (Violates `max: 1`).
* 3.4 :x: A folder containing only `.pdf` or `.epub` files but no `.zip` (Violates `.zip` `min: 1`).

## 4 Comic

* 4.1 :white_check_mark: A folder containing `issue1.zip`, `comic.nfo`, and `file_id.diz`.
* 4.2 :x: A folder missing `file_id.diz` (Violates `min: 1`).
* 4.3 :x: A folder containing `.cbr` or `.cbz` files but no `.zip` (Violates `.zip` `min: 1`).
* 4.4 :x: A folder with multiple `.nfo` files (Violates `max: 1`).

### 5 Education

* 5.1 :white_check_mark: A folder with `course.rar`, `course.r01`, `checksum.sfv`, and `info.nfo`.
* 5.2 :white_check_mark: A folder with `course.rar`, `course.r00`, `course.r01`, etc.
* 5.3 :x: A folder missing `.r???` volume files (Violates `min: 1`).
* 5.4 :x: A folder with two `.rar` files (Violates `max: 1`—only one main archive allowed).
* 5.5 :x: A folder missing the `.sfv` file (Violates `min: 1`).

### 6 Episode

* 6.1 :white_check_mark: Contains `show.rar`, `show.r01`, `show.sfv`, `show.nfo`, and a `Sample` folder containing one `.mkv` file.
* 6.2 :white_check_mark: Same as above, but the `Sample` folder contains one `.mp4` file.
* 6.3 :x: `Sample` folder is missing entirely (Violates directory `min: 1`).
* 6.4 :x: `Sample` folder is empty (Violates sample file `min: 1`).
* 6.5 :x: `Sample` folder contains *both* an `.mkv` and an `.mp4` (Violates sample file `max: 1`).
* 6.6 :x: `Sample` folder contains an `.avi` file (Violates pattern `.{mkv,mp4}`).
* 6.7 :x: Root folder is missing `.r???` files.

### 7 Game

* 7.1 :white_check_mark: A folder with `game.rar`, `game.r00`, `game.r01`, `checksum.sfv`, and `release.nfo`.
* 7.2 :x: A folder missing the `.nfo` file.
* 7.3 :x: A folder missing the `.rar` file (even if `.r00` exists).
* 7.4 :x: A folder containing two `.sfv` files (Violates `max: 1`).

### 8 Magazine

* 8.1 :white_check_mark: A folder containing `mag.zip`, `release.nfo`, and `file_id.diz`.
* 8.2 :x: A folder containing only `.pdf` files.
* 8.3 :x: A folder missing `file_id.diz`.
* 8.4 :x: A folder containing two `.nfo` files.

### 9 Movie

* 9.1 :white_check_mark: Contains `movie.rar`, `movie.r01`, `movie.sfv`, `movie.nfo`, and a `Sample` folder containing one `.mkv`.
* 9.2 :x: `Sample` folder contains a `.wmv` file (Violates extension restriction).
* 9.3 :x: The `Sample` folder exists but is empty.
* 9.4 :x: Missing the main `.rar` archive file.
* 9.5 :x: Missing the `.nfo` file.

### 10 Music

* 10.1 :white_check_mark: A standard album release with 12 `.mp3` files, 1 `.m3u`, 1 `.sfv`, and 1 `.nfo`.
* 10.2 :white_check_mark: A multi-disc release with 2 `.m3u` files (valid as no max is defined for playlists).
* 10.3 :x: A folder with `.flac` files only (Violates `.mp3` `min: 1`).
* 10.4 :x: A folder missing the `.sfv` file.
* 10.5 :x: A folder with multiple `.nfo` files.

### 11 Series

* 11.1 :white_check_mark: A folder containing 2 subdirectories (e.g., "Season 1", "Season 2").
* 11.2 :white_check_mark: A folder containing 5 subdirectories.
* 11.3 :x: A folder containing only 1 subdirectory (Violates `min: 2`).
* 11.4 :x: A folder containing only files and no directories.
