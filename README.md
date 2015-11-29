# JapanGeoJsonGo

47都道府県のGeoJsonを県・市・町・村・郡・区単位に分割したりTopoJsonを作るGo、shellです。

GeoJsonを細かい単位に分割して使いやすくするのが目的です。

# ダウンロード

- 元データを以下URLからダウンロードします。

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html


# コマンド

### ファイル設置

- dataディレクトリを作成してダウンロードしたデータを配置

~~~
% mkdir data
% ls data
KS-META-N03-15_150101.xml	N03-15_150101.prj		N03-15_150101.shx
N03-15_150101.dbf		N03-15_150101.shp		N03-15_150101.xml
~~~

### GeoJson作成

- japan.jsonという全国分のGeoJsonができます。

~~~
% make geojson
% ls -l japan.json
-rw-r--r--  1 tetsuya  staff  607312018 11 24 23:28 japan.json
~~~

### GeoJson分割

- geojsonディレクトリにgeojsonが出力されます。

~~~
% make geojson-split
~~~

##### GeoJson GoogleMap Test

- ブラウザでlocalhost:3000で確認。

- Ctrl + Cで終了。

~~~
% make geosample PREF=富山県 CITY=氷見市
~~~

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

- topojsonディレクトリにtopojsonが出力されます。（作成中）

~~~
% make topojson
~~~

##### TopoJson D3.js Test


### 市町村郡区シェイプ確認デモ(GoogleMap)

http://geojson.niiyz.com/


![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot2.png)


```
国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

GeoJson http://geojson.org/

TopoJson https://github.com/mbostock/topojson
```