usage:
	@echo ""
	@echo "Task                                         : Description"
	@echo "----------------------------------           : -------------------"
	@echo "geojson                                      : Convert SHAPE to GeoJson"
	@echo "geojson-split                                : Split Japan GeoJson to City"
	@echo "topojson                                     : Convert GeoJson to TopoJson"
	@echo "geosample PREF={富山県} CITY={氷見市}        : Test GeoJson on GoogleMap"
	@echo "toposample PREF={東京都} CITY={新宿区}       : Test TopoJson on D3js"
	@echo "clean                                        : Remove ALL GeoJson and TopoJson"

geojson:
	rm -rf ./japan.json
	ogr2ogr -f GeoJSON japan.json data/N03-15_150101.shp

geojson-split:
	@echo PLEASE WAIT....
	go run main.go

.PHONY: topojson
topojson:
	./script/geojson_to_topojson.sh

geosample:
	@echo ${PREF} ${CITY}
	@echo Access to http://localhost:3000
	go run ./sample/sample.go geojson $(PREF) $(CITY)

toposample:
	@echo ${PREF} ${CITY}
	@echo Access to http://localhost:3000
	go run ./sample/sample.go topojson $(PREF) $(CITY)
	
clean:
	rm -rf topojson
	rm -rf geojson
	