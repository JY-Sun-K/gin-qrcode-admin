function EuToUsCurrencyFormat(input) {
  return input.replace(/[,.]/g, function(x) {
    return x == "," ? "." : ",";
  });
}

$(document).ready(function() {
  //Only needed for the filename of export files.
  //Normally set in the title tag of your page.
  document.title = 'DataTable Excel';
  // DataTable initialisation
  $('#example').DataTable({
    "dom": '<"dt-buttons"Bf><"clear">lirtp',
    "paging": true,
    "autoWidth": true,
    "buttons": [{
      extend: 'excelHtml5',
      text: 'Excel',
      customize: function(xlsx) {
        var sheet = xlsx.xl.worksheets['sheet1.xml'];
        //All cells
        $('row c', sheet).attr('s', '25');
        //Second column
        $('row c:nth-child(2)', sheet).attr('s', '42');
        //First row
        $('row:first c', sheet).attr('s', '36');
        // One cell
        $('row c[r^="D6"]', sheet).attr('s', '32');
        // Loop over the cells in column `E` the amount column
        $('row c[r^="E"]', sheet).each(function() {
          if (parseFloat(EuToUsCurrencyFormat($('is t', this).text())) > 1500) {
            $(this).attr('s', '17');
          }
        });
        //All cells of row 10
        $('row c[r*="10"]', sheet).attr('s', '49');
        //Search all cells for a specific text
        $('row* c[r]', sheet).each(function() {
          if ($('is t', this).text().match(/(?:^|\b)(cover)(?=\b|$)/gmi)) {
            $(this).attr('s', '20');
          }
        });
      }
    }]
  });
});