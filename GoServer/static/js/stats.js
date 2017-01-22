$.getJSON( "/api/stats/brands/", function( data ) {
  var numData = [];
  $.each( data, function( key, val ) {
    numData.push({label:val.Name, value:val.Num});
  });
  console.log(numData);
  console.log(nv);
  nv.addGraph(function(){
    var chart = nv.models.pieChart()
    .x(function(d) { return d.label })
    .y(function(d) { return d.value })
    .width(100)
    .height(100);
    
    d3.select("#shareGraph")
    .datum(numData)
    .transition().duration(1200)
    .attr('width', 100)
    .attr('height', 100)
    .call(chart);
    
    return chart;
  })
});