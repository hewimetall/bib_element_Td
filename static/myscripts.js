const searchInputs = [...document.querySelectorAll('input')]
  .filter(({ id }) => id);


async function post(params) {
  const sp = new URLSearchParams(params);
  const url = `/api/searchs?${ sp.toString() }`;

  const response = await fetch(url);

  const json = await response.json();
  const data = json.result.map((item, index) => ({
    id: index,
    title: item.Title,
    author: item.author,
    publishDate: item.P_Date
  }));

  return data;
}

function drawTable(data) {
  console.dir(data)
  const table = new Tabulator("#example-table", {
    data,
    autoColumns:true,           //load row data from array
    layout:"fitColumns",      //fit columns to width of table
    responsiveLayout:"hide",  //hide columns that dont fit on the table
    tooltips:true,            //show tool tips on cells
    addRowPos:"top",          //when adding a new row, add it to the top of the table
    history:true,             //allow undo and redo actions on the table
    pagination:"local",       //paginate the data
    paginationSize:7,         //allow 7 rows per page of data
    movableColumns:true,      //allow column order to be changed
    resizableRows:true,       //allow row order to be changed
    initialSort:[             //set the initial sort order of the data
        {column:"name", dir:"asc"},
    ],
  });

  //table.element.addEve
}  

function getData() {
  const params = {};

  searchInputs.forEach(({ id, value }) => params[id] = value);
  
  post(params).then(data => drawTable(data));
}
