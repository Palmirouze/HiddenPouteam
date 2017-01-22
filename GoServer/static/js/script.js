$(document).ready(function() {
    $('#search').keyup(function () {
        $.get("/searchResult?q=" + $('#search').val(), function (data) {
            $("#result").html(data);
        });
    });
});