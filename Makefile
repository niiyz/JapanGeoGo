usage:
	@echo ""
	@echo "Task                                         : Description"
	@echo "----------------------------------           : -------------------"
	@echo "geojson                                      : Convert SHAPE to GeoJson"
	@echo "geojson-split                                : Split Japan GeoJson to City"
	@echo "topojson                                     : Convert GeoJson to TopoJson"
	@echo "gosample PREF=富山県 CITY=氷見市             : Test GeoJson on GoogleMap"

geojson:
	rm -rf ./japan.json
	ogr2ogr -f GeoJSON japan.json data/N03-15_150101.shp

geojson-split:
	@echo PLEASE WAIT....
	go run main.go

topojson:
	./script/geojson_to_topojson.sh

geosample:
	@echo ${PREF} ${CITY}
	go run ./sample/sample.go $(PREF) $(CITY)
	@echo Access to http://localhost:3000

toposample:
	@echo ${PREF} ${CITY}