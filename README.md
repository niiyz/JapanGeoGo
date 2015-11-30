# JapanGeoGo

47都道府県のGeoJsonを県・市・町・村・郡・区単位に分割するGo、TopoJsonを作るラッパーshellです。

GeoJsonを細かい単位に分割して使いやすくするのが目的です。

# 準備

Use MacOSX

~~~
% // ogr2ogr
% brew install gdal
% which ogr2ogr
% // topojson
% npm install -g topojson
% which topojson]
% // Golang
% brew install go
% go version
% go version go1.4.2 darwin/amd64
% // Clone
% git clone https://github.com/niiyz/JapanGeoGo.git
~~~

# ダウンロード

- 元データを以下URLからダウンロードします。

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

### ファイル設置

- dataディレクトリを作成してダウンロードしたデータを配置

~~~
% mkdir data
% ls data
KS-META-N03-15_150101.xml	N03-15_150101.prj		N03-15_150101.shx
N03-15_150101.dbf		N03-15_150101.shp		N03-15_150101.xml
~~~

# コマンド

### GeoJson作成

- GeoJson「japan.json」を出力。

~~~
% make geojson
% ls -l japan.json
-rw-r--r--  1 tetsuya  staff  607312018 11 24 23:28 japan.json
~~~

### GeoJson分割

- ディレクトリgeojsonを作成、配下にgeojson出力。

~~~
% make geojson-split
~~~

##### GeoJson GoogleMap Test

<img alt="screenshot_geojson_on_googlemap" width="500" src="https://github.com/niiyz/JapanGeoGo/blob/master/images/screenshot_geojson.png" />

- ブラウザでlocalhost:3000で確認。

- Ctrl + Cで終了。

~~~
% make geosample PREF=神奈川県 CITY=横浜市
~~~

~~~
% make geosample PREF=神奈川県 CITY=横浜市青葉区
~~~

~~~
% make geosample PREF=鹿児島県 CITY=薩摩郡
~~~

~~~
% make geosample PREF=鹿児島県 CITY=薩摩郡さつま町
~~~

~~~
% make geosample PREF=沖縄県
~~~

### TopoJson作成

- ディレクトリtopojsonを作成、配下にtopojson作成。

~~~
% make topojson
~~~

##### TopoJson D3.js Test

<img alt="screenshot_topojson_on_d3js" width="500" src="https://github.com/niiyz/JapanGeoGo/blob/master/images/screenshot_topojson.png" />

~~~
% make toposample PREF=富山県 CITY=高岡市
~~~

~~~
% make toposample PREF=北海道
~~~

### Reference

* 国土交通省 国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　: http://nlftp.mlit.go.jp
* GeoJson: http://geojson.org/
* TopoJson: https://github.com/mbostock/topojson
