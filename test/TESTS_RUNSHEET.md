# TESTS RUNSHEET

## 1 App

* **[1.1](validate/01_1/)** :white_check_mark: A folder containing `setup.zip`, `release.nfo`, and `file_id.diz`.

  ```bash
  ```

* **[1.2](validate/01_2/)** :white_check_mark: A folder containing `part1.zip`, `part2.zip`, `release.nfo`, and `file_id.diz` (Multiple zips allowed).

  ```bash
  ```

* **[1.3](validate/01_3/)** :x: A folder is missing `file_id.diz` (Violates `min: 1`).

  ```bash
  ```

* **[1.4](validate/01_4/)** :x: A folder containing two `.nfo` files (Violates `max: 1`).

  ```bash
  ```

* **[1.5](validate/01_5/)** :x: A folder with `installer.rar` instead of `.zip` (Violates pattern).

  ```bash
  ```

## 2 Audiobook

* **[2.1](validate/02_1/)** :white_check_mark: A folder with `playlist.m3u`, `checksum.sfv`, `book.nfo`, and 10 `.mp3` files.

  ```bash
  ```

* **[2.2](validate/02_2/)** :white_check_mark: A folder with multiple `.m3u` playlists (e.g., CD1, CD2) and multiple `.mp3` files.

  ```bash
  ```

* **[2.3](validate/02_3/)** :x: A folder missing the `.sfv` file (Violates `min: 1`).

  ```bash
  ```

* **[2.4](validate/02_4/)** :x: A folder containing `.m4b` audio files but no `.mp3` files (Violates `.mp3` `min: 1`).

  ```bash
  ```

* **[2.5](validate/02_5/)** :x: A folder with two `.nfo` files (Violates `max: 1`).

  ```bash
  ```

## 3 Book

* **[3.1](validate/03_1/)** :white_check_mark: A folder containing `novel.zip`, `info.nfo`, and `file_id.diz`.

  ```bash
  ```

* **[3.2](validate/03_2/)** :x: A folder missing the `.nfo` file (Violates `min: 1`).

  ```bash
  ```

* **[3.3](validate/03_3/)** :x: A folder containing two `.diz` files (Violates `max: 1`).

  ```bash
  ```

* **[3.4](validate/03_4/)** :x: A folder containing only `.pdf` or `.epub` files but no `.zip` (Violates `.zip` `min: 1`).

  ```bash
  ```

## 4 Comic

* **[4.1](validate/04_1/)** :white_check_mark: A folder containing `issue1.zip`, `comic.nfo`, and `file_id.diz`.

  ```bash
  ```

* **[4.2](validate/04_2/)** :x: A folder missing `file_id.diz` (Violates `min: 1`).

  ```bash
  ```

* **[4.3](validate/04_3/)** :x: A folder containing `.cbr` or `.cbz` files but no `.zip` (Violates `.zip` `min: 1`).

  ```bash
  ```

* **[4.4](validate/04_4/)** :x: A folder with multiple `.nfo` files (Violates `max: 1`).

  ```bash
  ```

## 5 Education

* **[5.1](validate/05_1/)** :white_check_mark: A folder with `course.rar`, `course.r01`, `checksum.sfv`, and `info.nfo`.

  ```bash
  ```

* **[5.2](validate/05_2/)** :white_check_mark: A folder with `course.rar`, `course.r00`, `course.r01` - `course.r99`, etc.

  ```bash
  ```

* **[5.3](validate/05_3/)** :x: A folder missing `.r??` volume files (Violates `min: 1`).

  ```bash
  ```

* **[5.4](validate/05_4/)** :x: A folder with two `.rar` files (Violates `max: 1` - only one main archive allowed).

  ```bash
  ```

* **[5.5](validate/05_5/)** :x: A folder missing the `.sfv` file (Violates `min: 1`).

  ```bash
  ```

## 6 Episode

* **[6.1](validate/06_1/)** :white_check_mark: Contains `show.rar`, `show.r01`, `show.sfv`, `show.nfo`, and a `Sample` folder containing one `.mkv` file.

  ```bash
  ```

* **[6.2](validate/06_2/)** :white_check_mark: Same as above, but the `Sample` folder contains one `.mp4` file.

  ```bash
  ```

* **[6.3](validate/06_3/)** :x: `Sample` folder is missing entirely (Violates directory `min: 1`).

  ```bash
  ```

* **[6.4](validate/06_4/)** :x: `Sample` folder is empty (Violates sample file `min: 1`).

  ```bash
  ```

* **[6.5](validate/06_5/)** :x: `Sample` folder contains both `.mkv` and an `.mp4` (Violates sample file `max: 1`).

  ```bash
  ```

* **[6.6](validate/06_6/)** :x: `Sample` folder contains an `.rar` file (Violates pattern `.{mkv,mp4}`).

  ```bash
  ```

* **[6.7](validate/06_7/)** :x: Root folder is missing `.r??` files.

  ```bash
  ```

## 7 Game

* **[7.1](validate/07_1/)** :white_check_mark: A folder with `game.rar`, `game.r00`, `game.r01`, `checksum.sfv`, and `release.nfo`.

  ```bash
  ```

* **[7.2](validate/07_2/)** :x: A folder missing the `.nfo` file.

  ```bash
  ```

* **[7.3](validate/07_3/)** :x: A folder missing the `.rar` file (even if `.r00` exists).

  ```bash
  ```

* **[7.4](validate/07_4/)** :x: A folder containing two `.sfv` files (Violates `max: 1`).

  ```bash
  ```

## 8 Magazine

* **[8.1](validate/08_1/)** :white_check_mark: A folder containing `mag.zip`, `release.nfo`, and `file_id.diz`.

  ```bash
  ```

* **[8.2](validate/08_2/)** :x: A folder containing only `.pdf` files.

  ```bash
  ```

* **[8.3](validate/08_3/)** :x: A folder missing `file_id.diz`.

  ```bash
  ```

* **[8.4](validate/08_4/)** :x: A folder containing two `.nfo` files.

  ```bash
  ```

## 9 Movie

* **[9.1](validate/09_1/)** :white_check_mark: Contains `movie.rar`, `movie.r01`, `movie.sfv`, `movie.nfo`, and a `Sample` folder containing one `.mkv`.

  ```bash
  ```

* **[9.2](validate/09_2/)** :x: `Sample` folder contains a `.wmv` file (Violates extension restriction).

  ```bash
  ```

* **[9.3](validate/09_3/)** :x: The `Sample` folder exists but is empty.

  ```bash
  ```

* **[9.4](validate/09_4/)** :x: Missing the main `.rar` archive file.

  ```bash
  ```

* **[9.5](validate/09_5/)** :x: Missing the `.nfo` file.

  ```bash
  ```

## 10 Music

* **[10.1](validate/10_1/)** :white_check_mark: A standard album release with 12x `.mp3` files, 1x `.m3u`, 1x `.sfv`, and 1x `.nfo`.

  ```bash
  ```

* **[10.2](validate/10_2/)** :white_check_mark: A multi-disc release with 2 `.m3u` files (valid as no max is defined for playlists).

  ```bash
  ```

* **[10.3](validate/10_3/)** :x: A folder with `.flac` files only.

  ```bash
  ```

* **[10.4](validate/10_4/)** :x: A folder missing the `.sfv` file.

  ```bash
  ```

* **[10.5](validate/10_5/)** :x: A folder with multiple `.nfo` files.

  ```bash
  ```

## 11 Series

* **[11.1](validate/11_1/)** :white_check_mark: A folder containing 2 subdirectories (e.g., "Season 1", "Season 2").

  ```bash
  ```

* **[11.2](validate/11_2/)** :white_check_mark: A folder containing 5 subdirectories.

  ```bash
  ```

* **[11.3](validate/11_3/)** :x: A folder containing only 1 subdirectory (Violates `min: 2`).

  ```bash
  ```

* **[11.4](validate/11_4/)** :x: A folder containing only files and no directories.

  ```bash
  ```
