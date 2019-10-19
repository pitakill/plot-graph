(async () => {
  try {
    // Initiate new instance of sigma
    const s = new sigma('container');

    // Get info
    const r = await fetch(url);
    const data = await r.json();

    // Add Nodes
    data.forEach(e => {
      s.graph.addNode({
        id: e.name,
        label: e.name,
        color: colorNodes[e.associations.length],
        size: 1,
        x: Math.random(),
        y: Math.random(),
        details: e.azDetails,
      });
    });

    // Add Edges
    data.forEach(e => {
      e.associations.forEach(n => {
        console.log()
        s.graph.addEdge({
          id: `${e.name}=>${n.name}`,
          source: e.name,
          target: n.name,
        });
      });
    });

  // Initiate drag nodes
  const dragListener = new sigma.plugins.dragNodes(s, s.renderers[0]);

  // Print in info the info of the node
  function showInfo(e) {
    const { children: [title, pre] } = document.getElementById("info");
    const data = JSON.stringify(e.data.node.details, null, 2);
    title.innerText = e.data.node.id;
    pre.innerText = data ? data : 'No addition data';
  };

  // Event listener
  s.bind('overNode', showInfo);
  dragListener.bind('drag', showInfo);
  //s.bind('outNode', function(e) {
    //const { children: [title, p] } = document.getElementById("info");
    //title.innerText = "";
    //p.innerText = "";
  //});

  // Display graph on browser
  s.refresh();
  } catch (e) {
    console.error(e);
  }
})();
