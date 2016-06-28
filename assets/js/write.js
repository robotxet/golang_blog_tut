$(function () {
    $("#content").bind("import change", function () {
        $.post("/gethtml", {md: $("#content").val()}, function (response) {
            $("#md_html").html(response.html)
        })
    })
})