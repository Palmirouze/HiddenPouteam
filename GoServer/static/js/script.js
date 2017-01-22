$(document).ready(function() {
    $('#search').keyup(function () {
        $.get("/productSearchResult?q=" + $('#search').val(), function (data) {
            $("#result").html(data);
        });
    });
});