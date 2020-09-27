async function post(params) {
  const response = await fetch('/api', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(params)
  });

  const json = await response.json();

  const data = json.result.map((item, index) => ({
    id: index,
    title: item.Title,
    author: item.author,
    publishDate: item.P_Date
  }));

  return data;
}
