/**
 * Created by Tim on 2017-01-21.
 */

<script>
$('#search').keyup(function(){
    console.log("Key press registered: " + $('#search').val());
    $.get( "/searchResult?q="+$('#search').val(), function( data ) {
        $("#result").html(data);
    });
});
</script>