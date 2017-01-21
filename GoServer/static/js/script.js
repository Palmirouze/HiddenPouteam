$(document).ready(function() {
    $('#search').keyup(function () {
        console.log("Key press registered: " + $('#search').val());
        $.get("/searchResult?q=" + $('#search').val(), function (data) {
            $("#result").html(data);
        });
    });
});