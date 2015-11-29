# JapanCityGeoJson

47都道府県の県・市・町・村・郡・区の形を作るためのGeoJsonデータ、TopoJsonデータです。

分割して必要な地域のファイルのみにするのが目的です。

# コマンド

### dataディレクトリにファイル設置

~~~
% ls data
KS-META-N03-15_150101.xml	N03-15_150101.prj		N03-15_150101.shx
N03-15_150101.dbf		N03-15_150101.shp		N03-15_150101.xml
~~~

### GeoJson作成

~~~
% make geojson
~~~

### GeoJson分割

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
make geosample PREF=神奈川県 CITY=横浜市
~~~

~~~
make geosample PREF=神奈川県 CITY=横浜市青葉区
~~~

~~~
% make geosample PREF=鹿児島県 CITY=薩摩郡さつま町
~~~

### TopoJson作成

~~~
% make topojson
~~~

##### TopoJson D3.js Test


### 市町村郡区シェイプ確認デモ(GoogleMap)

http://geojson.niiyz.com/


![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot2.png)


```
国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

GeoJson http://geojson.org/

TopoJson https://github.com/mbostock/topojson
```