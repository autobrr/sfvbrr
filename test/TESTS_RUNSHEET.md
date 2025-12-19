# TESTS RUNSHEET

## 1 App

* **[1.1](validate/01_1/)** :white_check_mark: A folder containing `setup.zip`, `release.nfo`, and `file_id.diz`.

  ```bash
  Validating Release:
    Folder:       01_1/App.Pro.v1.1.1.Linux.RPM.ARM64.Incl.Keymaker-GRP
    Category:     app

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[1.2](validate/01_2/)** :white_check_mark: A folder containing `part1.zip`, `part2.zip`, `release.nfo`, and `file_id.diz` (Multiple zips allowed).

  ```bash
  Validating Release:
    Folder:       01_2/Corporation.Program.v1.0.0.0.x64.Multilingual.Incl.Keymaker-GRP
    Category:     app

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[1.3](validate/01_3/)** :x: A folder is missing `file_id.diz` (Violates `min: 1`).

  ```bash
  Validating Release:
    Folder:       01_3/APPLICATION.TEXT.TEXT.V2025.2-GRP
    Category:     app

  Rule Validation:
    ✗ file_id.diz - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[1.4](validate/01_4/)** :x: A folder containing two `.nfo` files (Violates `max: 1`).

  ```bash
  Validating Release:
    Folder:       01_4/AppApp.v10.00.MacOS.UB.Incl.Keygen-GRP
    Category:     app

  Rule Validation:
    ✗ *.nfo (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[1.5](validate/01_5/)** :x: A folder with `installer.rar` instead of `.zip` (Violates pattern).

  ```bash
  Validating Release:
    Folder:       01_5/Software.Tool.Compiler.v2025.20.x86.Incl.Keygen-GRP
    Category:     app

  Rule Validation:
    ✗ *.zip - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Unexpected Files/Directories:
    ✗ installer.rar

  Errors:
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[1.6](validate/01_6/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       01_6/App.Pro.v1.1.1.Linux.RPM.ARM64.Incl.Keymaker-GRP
    Category:     app

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.txt

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 2 Audiobook

* **[2.1](validate/02_1/)** :white_check_mark: A folder with `playlist.m3u`, `checksum.sfv`, `book.nfo`, and 10 `.mp3` files.

  ```bash
  Validating Release:
    Folder:       02_1/First_Last_-_Title-AUDiOBOOK-WEB-EN-2025-GRP
    Category:     audiobook

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[2.2](validate/02_2/)** :white_check_mark: A folder with multiple `.m3u` playlists (e.g., CD1, CD2) and multiple `.mp3` files.

  ```bash
  Validating Release:
    Folder:       02_2/Artist-Title-AUDiOBOOK-WEB-EN-2025-GRP
    Category:     audiobook

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[2.3](validate/02_3/)** :x: A folder missing the `.sfv` file (Violates `min: 1`).

  ```bash
  Validating Release:
    Folder:       02_3/First_Last_-_1_2_3_4_5_6-AUDiOBOOK-WEB-SE-2025-GRP
    Category:     audiobook

  Rule Validation:
    ✗ *.sfv - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[2.4](validate/02_4/)** :x: A folder containing `.m4b` audio files but no `.mp3` files (Violates `.mp3` `min: 1`).

  ```bash
  Validating Release:
    Folder:       02_4/First_Last_-_Title-AUDiOBOOK-WEB-EN-2025-GRP
    Category:     audiobook

  Rule Validation:
    ✗ *.mp3 - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Unexpected Files/Directories:
    ✗ 01-track.m4b

  Errors:
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[2.5](validate/02_5/)** :x: A folder with two `.nfo` files (Violates `max: 1`).

  ```bash
  Validating Release:
    Folder:       02_5/First_Last_-_Title-AUDiOBOOK-WEB-EN-2025-GRP
    Category:     audiobook

  Rule Validation:
    ✗ *.nfo (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[2.6](validate/02_6/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       02_6/First_Last_-_Title-AUDiOBOOK-WEB-EN-2025-GRP
    Category:     audiobook

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.txt

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 3 Book

* **[3.1](validate/03_1/)** :white_check_mark: A folder containing `novel.zip`, `info.nfo`, and `file_id.diz`.

  ```bash
  Validating Release:
    Folder:       03_1/First.Last.The.Title.2025.RETAiL.ePub.eBook-GRP
    Category:     book

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[3.2](validate/03_2/)** :x: A folder missing the `.nfo` file (Violates `min: 1`).

  ```bash
  Validating Release:
    Folder:       03_2/Publishing.Title.Second.Edition.Artbook.2025.Scan.eBook-GRP
    Category:     book

  Rule Validation:
    ✗ *.nfo - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[3.3](validate/03_3/)** :x: A folder containing two `.diz` files (Violates `max: 1`).

  ```bash
  Validating Release:
    Folder:       03_3/Publisher.-.Word.Word.Word.3rd.Edition.2025.Retail.EPUB.eBook-GRP
    Category:     book

  Rule Validation:
    ✗ *.diz (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[3.4](validate/03_4/)** :x: A folder containing only `.pdf` or `.epub` files but no `.zip` (Violates `.zip` `min: 1`).

  ```bash
  Validating Release:
    Folder:       03_4/Books.-.La.La.La.La.La.La.La.La.La.2025.Retail.eBook-GRP
    Category:     book

  Rule Validation:
    ✗ *.nfo - found 0 matches, but minimum required is 1
    ✗ file_id.diz - found 0 matches, but minimum required is 1
    ✗ *.zip - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    1
    Invalid rules:  3

  Unexpected Files/Directories:
    ✗ file1.pdf
    ✗ file2.epub

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
    found 2 unexpected file(s)/directory(ies)
  ```

* **[3.5](validate/03_5/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       03_5/First.Last.The.Title.2025.RETAiL.ePub.eBook-GRP
    Category:     book

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ Some other folder

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 4 Comic

* **[4.1](validate/04_1/)** :white_check_mark: A folder containing `issue1.zip`, `comic.nfo`, and `file_id.diz`.

  ```bash
  Validating Release:
    Folder:       04_1/Comics.-.Title.Vol.01.2025.Retail.Comic.eBook-GRP
    Category:     comic

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[4.2](validate/04_2/)** :x: A folder missing `file_id.diz` (Violates `min: 1`).

  ```bash
  Validating Release:
    Folder:       04_2/Comics.-.Title.Vol.02.2025.Retail.Comic.eBook-GRP
    Category:     comic

  Rule Validation:
    ✗ file_id.diz - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[4.3](validate/04_3/)** :x: A folder containing `.cbr` or `.cbz` files but no `.zip` (Violates `.zip` `min: 1`).

  ```bash
  Validating Release:
    Folder:       04_3/Comics.-.Title.Vol.03.2025.Retail.Comic.eBook-GRP
    Category:     comic

  Rule Validation:
    ✗ *.zip - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Unexpected Files/Directories:
    ✗ issue1.cbr

  Errors:
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[4.4](validate/04_4/)** :x: A folder with multiple `.nfo` files (Violates `max: 1`).

  ```bash
  Validating Release:
    Folder:       04_4/Comics.-.Title.Vol.04.2025.Retail.Comic.eBook-GRP
    Category:     comic

  Rule Validation:
    ✗ *.nfo (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[4.5](validate/04_5/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       04_5/Comics.-.Title.Vol.01.2025.Retail.Comic.eBook-GRP
    Category:     comic

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.txt

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 5 Education

* **[5.1](validate/05_1/)** :white_check_mark: A folder with `course.rar`, `course.r01`, `checksum.sfv`, and `info.nfo`.

  ```bash
  Validating Release:
    Folder:       05_1/Learning.-.Topic.UPDATED.November.2025.BOOKWARE-GRP
    Category:     education

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[5.2](validate/05_2/)** :white_check_mark: A folder with `course.rar`, `course.r00`, `course.r01` - `course.r99`, etc.

  ```bash
  Validating Release:
    Folder:       05_2/OREILLY_CLOUD-GRP
    Category:     education

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[5.3](validate/05_3/)** :x: A folder missing `.r??` volume files (Violates `min: 1`).

  ```bash
  Validating Release:
    Folder:       05_3/WHATEVER_TUTORIAL-ISO
    Category:     education

  Rule Validation:
    ✗ *.rar - found 0 matches, but minimum required is 1
    ✗ .*\.r\d{2}$ - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    2
    Invalid rules:  2

  Unexpected Files/Directories:
    ✗ tutorial.bin

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[5.4](validate/05_4/)** :x: A folder with two `.rar` files (Violates `max: 1` - only one main archive allowed).

  ```bash
  Validating Release:
    Folder:       05_4/Udemy.The.Best.Seminar.TUTORiAL.GERMAN-GRP
    Category:     education

  Rule Validation:
    ✗ *.rar (found 2) - found 2 matches, but maximum allowed is 1
    ✗ .*\.r\d{2}$ - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    2
    Invalid rules:  2

  Errors:
    found 2 matches, but maximum allowed is 1
    found 0 matches, but minimum required is 1
  ```

* **[5.5](validate/05_5/)** :x: A folder missing the `.sfv` file (Violates `min: 1`).

  ```bash
  Validating Release:
    Folder:       05_5/Udemy.The.Best.Seminar.BOOKWARE.SWEDISH-GRP
    Category:     education

  Rule Validation:
    ✗ *.sfv - found 0 matches, but minimum required is 1
    ✗ .*\.r\d{2}$ - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    2
    Invalid rules:  2

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
  ```

* **[5.6](validate/05_6/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       05_6/Learning.-.Topic.UPDATED.November.2025.BOOKWARE-GRP
    Category:     education

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ Whatever

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 6 Episode

* **[6.1](validate/06_1/)** :white_check_mark: Contains `show.rar`, `show.r01`, `show.sfv`, `show.nfo`, and a `Sample` folder containing one `.mkv` file.

  ```bash
  Validating Release:
    Folder:       06_1/Show.S01E06.1080p.WEB.h264-GRP
    Category:     episode

  Rule Validation:

  Summary:
    Valid rules:    6
  ```

* **[6.2](validate/06_2/)** :white_check_mark: Same as above, but the `Sample` folder contains one `.mp4` file.

  ```bash
  Validating Release:
    Folder:       06_2/TVShow.S10.E999.720p.BluRay.x264-GRP
    Category:     episode

  Rule Validation:

  Summary:
    Valid rules:    6
  ```

* **[6.3](validate/06_3/)** :x: `Sample` folder is missing entirely (Violates directory `min: 1`).

  ```bash
  Validating Release:
    Folder:       06_3/GreatShow.S09E01.1080p.WEB.x264-GRP
    Category:     episode

  Rule Validation:
    ✗ Sample - found 0 matches, but minimum required is 1
    ✗ Sample/*.{mkv,mp4} - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    4
    Invalid rules:  2

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
  ```

* **[6.4](validate/06_4/)** :x: `Sample` folder is empty (Violates sample file `min: 1`).

  ```bash
  Validating Release:
    Folder:       06_4/Top.Show.S01E01.2025.S01E10.DV.2160p.WEB.h265-GRP
    Category:     episode

  Rule Validation:
    ✗ Sample - found 0 matches, but minimum required is 1
    ✗ Sample/*.{mkv,mp4} - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    4
    Invalid rules:  2

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
  ```

* **[6.5](validate/06_5/)** :x: `Sample` folder contains both `.mkv` and an `.mp4` (Violates sample file `max: 1`).

  ```bash
  Validating Release:
    Folder:       06_5/Episode.E01.MULTi.1080p.WEB.x264-GRP
    Category:     episode

  Rule Validation:
    ✗ Sample/*.{mkv,mp4} (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    5
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[6.6](validate/06_6/)** :x: `Sample` folder contains an `.rar` file (Violates pattern `.{mkv,mp4}`).

  ```bash
  Validating Release:
    Folder:       06_6/Episode.S01E99.2025.HDR.2160p.WEB.h265-GRP
    Category:     episode

  Rule Validation:
    ✗ Sample/*.{mkv,mp4} - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    5
    Invalid rules:  1

  Unexpected Files/Directories:
    ✗ Sample/show.rar

  Errors:
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[6.7](validate/06_7/)** :x: Root folder is missing `.r??` files.

  ```bash
  Validating Release:
    Folder:       06_7/EPISODE.25.01.01.XXX.1080p.MP4-GRP
    Category:     episode

  Rule Validation:
    ✗ .*\.r\d{2}$ - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    5
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[6.8](validate/06_8/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       06_8/Show.S01E06.1080p.WEB.h264-GRP
    Category:     episode

  Rule Validation:

  Summary:
    Valid rules:    6

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.txt

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

* **[6.9](validate/06_9/)** :white_check_mark: A valid release folder with one additional JPEG file in the Proof folder.

  ```bash
  Validating Release:
    Folder:       06_9/Show.S01E09.1080p.BluRay.x264-GRP
    Category:     episode

  Rule Validation:

  Summary:
    Valid rules: 8
  ```

## 7 Game

* **[7.1](validate/07_1/)** :white_check_mark: A folder with `game.rar`, `game.r00`, `game.r01`, `checksum.sfv`, and `release.nfo`.

  ```bash
  Validating Release:
    Folder:       07_1/Great.Game.Udate.v1.2.34.567890.incl.DLC-GRP
    Category:     game

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[7.2](validate/07_2/)** :x: A folder missing the `.nfo` file.

  ```bash
  Validating Release:
    Folder:       07_2/Game1.v2.0.0.0.GOG-GRP
    Category:     game

  Rule Validation:
    ✗ *.nfo - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[7.3](validate/07_3/)** :x: A folder missing the `.rar` file (even if `.r00` exists).

  ```bash
  Validating Release:
    Folder:       07_3/Game2.GoG.Classic-GRP
    Category:     game

  Rule Validation:
    ✗ *.rar - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[7.4](validate/07_4/)** :x: A folder containing two `.sfv` files (Violates `max: 1`).

  ```bash
  Validating Release:
    Folder:       07_4/Game_Boy_Nintendo_Switch_Online_Update_v1.0.0_INTERNAL_JPN_NSW-GRP
    Category:     game

  Rule Validation:
    ✗ *.sfv (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[7.5](validate/07_5/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       07_5/Great.Game.Udate.v1.2.34.567890.incl.DLC-GRP
    Category:     game

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.bin

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 8 Magazine

* **[8.1](validate/08_1/)** :white_check_mark: A folder containing `mag.zip`, `release.nfo`, and `file_id.diz`.

  ```bash
  Validating Release:
    Folder:       08_1/Title.No.100.2025.GERMAN.HYBRID.MAGAZINE.eBook-GRP
    Category:     magazine

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[8.2](validate/08_2/)** :x: A folder containing only `.pdf` files.

  ```bash
  Validating Release:
    Folder:       08_2/Title.No.1.2025.FiNNiSH.HYBRiD.MAGAZiNE.eBook-GRP
    Category:     magazine

  Rule Validation:
    ✗ *.nfo - found 0 matches, but minimum required is 1
    ✗ file_id.diz - found 0 matches, but minimum required is 1
    ✗ *.zip - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    1
    Invalid rules:  3

  Unexpected Files/Directories:
    ✗ mag.pdf

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[8.3](validate/08_3/)** :x: A folder missing `file_id.diz`.

  ```bash
  Validating Release:
    Folder:       08_3/Title.No.1.2025.NORWEGiAN.HYBRiD.MAGAZiNE.eBook-GRP
    Category:     magazine

  Rule Validation:
    ✗ file_id.diz - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[8.4](validate/08_4/)** :x: A folder containing two `.nfo` files.

  ```bash
  Validating Release:
    Folder:       08_4/Mag.No.123.2025.THAi.HYBRiD.MAGAZiNE.eBook-GRP
    Category:     magazine

  Rule Validation:
    ✗ *.nfo (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[8.5](validate/08_5/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       08_5/Title.No.100.2025.GERMAN.HYBRID.MAGAZINE.eBook-GRP
    Category:     magazine

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ Completely_random_folder

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 9 Movie

* **[9.1](validate/09_1/)** :white_check_mark: Contains `movie.rar`, `movie.r01`, `movie.sfv`, `movie.nfo`, and a `Sample` folder containing one `.mkv`.

  ```bash
  Validating Release:
    Folder:       09_1/The.Movie.2025.1080P.BLURAY.X264-GRP
    Category:     movie

  Rule Validation:

  Summary:
    Valid rules:    6
  ```

* **[9.2](validate/09_2/)** :x: `Sample` folder contains a `.wmv` file (Violates extension restriction).

  ```bash
  Validating Release:
    Folder:       09_2/Movie.1900.COMPLETE.BLURAY-GRP
    Category:     movie

  Rule Validation:
    ✗ Sample/*.{mkv,mp4} - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    5
    Invalid rules:  1

  Unexpected Files/Directories:
    ✗ Sample/sample.wmv

  Errors:
    found 0 matches, but minimum required is 1
    found 1 unexpected file(s)/directory(ies)
  ```

* **[9.3](validate/09_3/)** :x: The `Sample` folder exists but is empty.

  ```bash
  Validating Release:
    Folder:       09_3/Great.Movie.of.2025.1080p.BluRay.H264-GRP
    Category:     movie

  Rule Validation:
    ✗ Sample - found 0 matches, but minimum required is 1
    ✗ Sample/*.{mkv,mp4} - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    4
    Invalid rules:  2

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
  ```

* **[9.4](validate/09_4/)** :x: Missing the main `.rar` archive file.

  ```bash
  Validating Release:
    Folder:       09_4/Another.Naughty.Movie.1900.XXX.BDRIP.X264-GRP
    Category:     movie

  Rule Validation:
    ✗ *.rar - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    5
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[9.5](validate/09_5/)** :x: Missing the `.nfo` file.

  ```bash
  Validating Release:
    Folder:       09_5/Movie.2025.SUBBED.DV.HDR.2160p.WEB.h265-GRP
    Category:     movie

  Rule Validation:
    ✗ *.nfo - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    5
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[9.6](validate/09_6/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       09_6/The.Movie.2025.1080P.BLURAY.X264-GRP
    Category:     movie

  Rule Validation:

  Summary:
    Valid rules:    6

  Unexpected Files/Directories:
    ✗ movie.srt

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

## 10 Music

* **[10.1](validate/10_1/)** :white_check_mark: A standard album release with 12x `.mp3` files, 1x `.m3u`, 1x `.sfv`, and 1x `.nfo`.

  ```bash
  Validating Release:
    Folder:       10_1/First_Last_-_Title-WEB-CZ-2025-GRP
    Category:     music

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[10.2](validate/10_2/)** :white_check_mark: A multi-disc release with 2 `.m3u` files (valid as no max is defined for playlists).

  ```bash
  Validating Release:
    Folder:       10_2/Artist-Title-WEB-2025-GRP
    Category:     music

  Rule Validation:

  Summary:
    Valid rules:    4
  ```

* **[10.3](validate/10_3/)** :x: A folder with `.flac` files only.

  ```bash
  Validating Release:
    Folder:       10_3/Artist-Album-DELUXE-24BIT-48KHZ-WEB-FLAC-2025-GRP
    Category:     music

  Rule Validation:
    ✗ *.m3u - found 0 matches, but minimum required is 1
    ✗ *.sfv - found 0 matches, but minimum required is 1
    ✗ *.nfo - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    1
    Invalid rules:  3

  Errors:
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
    found 0 matches, but minimum required is 1
  ```

* **[10.4](validate/10_4/)** :x: A folder missing the `.sfv` file.

  ```bash
  Validating Release:
    Folder:       10_4/First_Last_-_Title-SINGLE-WEB-EN-2025-GRP
    Category:     music

  Rule Validation:
    ✗ *.sfv - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[10.5](validate/10_5/)** :x: A folder with multiple `.nfo` files.

  ```bash
  Validating Release:
    Folder:       10_5/BAND-ALBUM-(CAT0001)-CD-FLAC-2025-GRP
    Category:     music

  Rule Validation:
    ✗ *.nfo (found 2) - found 2 matches, but maximum allowed is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 2 matches, but maximum allowed is 1
  ```

* **[10.6](validate/10_6/)** :x: A folder with no `.nfo` files.

  ```bash
  Validating Release:
    Folder:       10_6/First_Last-Title-WEB-SK-2025-GRP
    Category:     music

  Rule Validation:
    ✗ *.nfo - found 0 matches, but minimum required is 1

  Summary:
    Valid rules:    3
    Invalid rules:  1

  Errors:
    found 0 matches, but minimum required is 1
  ```

* **[10.7](validate/10_7/)** :x: A valid release folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       10_7/First_Last_-_Title-WEB-CZ-2025-GRP
    Category:     music

  Rule Validation:

  Summary:
    Valid rules:    4

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.bmp

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```

* **[10.8](validate/10_8/)** :x: A valid release folder with one additional unassociated file

  ```bash
  Validating Release:
    Folder:       10_8/First_Last_-_Title-WEB-CZ-2025-GRP
    Category:     music

  Rule Validation:

  Summary:
    Valid rules: 5
  ```


## 11 Series

* **[11.1](validate/11_1/)** :white_check_mark: A folder containing 2 subdirectories (e.g., "Season 1", "Season 2").

  ```bash
  Validating Release:
    Folder:       11_1/Show.Season.S01.COMPLETE.HDTV.x264-GRP
    Category:     series

  Rule Validation:

  Summary:
    Valid rules:    1
  ```

* **[11.2](validate/11_2/)** :white_check_mark: A folder containing 5 subdirectories.

  ```bash
  Validating Release:
    Folder:       11_2/Show.S01.1080p.WEB.H264-GRP
    Category:     series

  Rule Validation:

  Summary:
    Valid rules:    1
  ```

* **[11.3](validate/11_3/)** :x: A folder containing only 1 subdirectory (Violates `min: 2`).

  ```bash
  Validating Release:
    Folder:       11_3/Super.Duper.Series.S02.1080p.BluRay.H264-GRP
    Category:     series

  Rule Validation:
    ✗ * (found 1) - found 1 matches, but minimum required is 2

  Summary:
    Valid rules:    0
    Invalid rules:  1

  Errors:
    found 1 matches, but minimum required is 2
  ```

* **[11.4](validate/11_4/)** :x: A folder containing only files and no directories.

  ```bash
  Validating Release:
    Folder:       11_4/The.Adventures.Of.Nobody.S32.720p.BluRay.x264-GRP
    Category:     series

  Rule Validation:
    ✗ * - found 0 matches, but minimum required is 2

  Summary:
    Valid rules:    0
    Invalid rules:  1

  Unexpected Files/Directories:
    ✗ completely_unrelated_file.txt
    ✗ imdb.nfo

  Errors:
    found 0 matches, but minimum required is 2
    found 2 unexpected file(s)/directory(ies)
  ```

* **[11.5](validate/11_5/)** :x: A valid folder with one additional unassociated file (Violates `deny_unexpected`).

  ```bash
  Validating Release:
    Folder:       11_5/Show.Season.S02.COMPLETE.HDTV.x264-GRP
    Category:     series

  Rule Validation:

  Summary:
    Valid rules:    1

  Unexpected Files/Directories:
    ✗ unassociated.log

  Errors:
    found 1 unexpected file(s)/directory(ies)
  ```
