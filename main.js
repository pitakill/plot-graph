// Websocket
let ws = new WebSocket(URLConnections);
// Initiate new instance of sigma
sigma.classes.graph.addMethod('nodeExists', function(node) {
  return !!this.nodesArray.find(e => e.id === node)
});
sigma.classes.graph.addMethod('edgeExists', function(edge) {
  return !!this.edgesArray.find(e => e.id === edge)
});
sigma.classes.graph.addMethod('nodesExist', function(nodeA, nodeB) {
  return this.nodeExists(nodeA) && this.nodeExists(nodeB)
});
const connections = new sigma('connections');
const resources = new sigma('resources');

// Plot
const plotResources = (data, element) => {
  // Add Nodes
  data.forEach(e => {
    // Break the execution if node already exists
    if (element.graph.nodeExists(e.name)) return
    element.graph.addNode({
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
      const id = `${e.name}=>${n.name}`;
      // Break the execution if edge already exists
      if (element.graph.edgeExists(id)) return
      // Break the execution if nodes do not exist
      if (!element.graph.nodesExist(e.name, n.name)) return

      element.graph.addEdge({
        id,
        source: e.name,
        target: n.name,
      });
    });
  });

  // Initiate drag nodes
  const dragListener = new sigma.plugins.dragNodes(element, element.renderers[0]);

  // Print in info the info of the node
  function showInfo(e) {
    const { children: [title, pre] } = document.getElementById("info");
    const data = JSON.stringify(e.data.node.details, null, 2);
    title.innerText = e.data.node.id;
    pre.innerText = data ? data : 'No addition data';
  };

  // Event listener
  element.bind('overNode', showInfo);
  dragListener.bind('drag', showInfo);

  // Display graph on browser
  element.refresh();
};
// Connections to plot
let conns = {
  nodes: [],
  edges: [],
};
const plotConnections = (data, element) => {
  // Add Nodes
  data.nodes.forEach(({ id, details }) => {
    // Break the execution if node already exists
    if (element.graph.nodeExists(id)) return
    element.graph.addNode({
      id,
      label: details.hostname ? details.hostname : id,
      color: colorNodes[0],
      size: 1,
      x: Math.random(),
      y: Math.random(),
      details,
    });
  });

  // Add Edges
  data.edges.forEach(({ source, destination }) => {
    const id = `${source}=>${destination}`;
    // Break the execution if edge already exists
    if (element.graph.edgeExists(id)) return
    // Break the execution if nodes do not exist
    if (!element.graph.nodesExist(source, destination)) return

    element.graph.addEdge({
      id,
      source,
      target: destination,
    });
  });

  // Initiate drag nodes
  const dragListener = new sigma.plugins.dragNodes(element, element.renderers[0]);

  // Print in info the info of the node
  function showInfo(e) {
    const { children: [title, pre] } = document.getElementById("info");
    const data = JSON.stringify(e.data.node.details, null, 2);
    title.innerText = e.data.node.label;
    pre.innerText = data ? data : 'No addition data';
  };

  // Event listener
  element.bind('overNode', showInfo);
  dragListener.bind('drag', showInfo);

  // Display graph on browser
  element.refresh();
};

// Get Connections from Socket
ws.addEventListener('message', e => {
  const data = JSON.parse(e.data);
  conns[`${data.type}s`].push(data);
  plotConnections(conns, connections);
});

// Get data from URL
(async () => {
  try {
    // Get info
    const r = await fetch(url);
    const data = await r.json();

    plotResources(data, resources);
  } catch (e) {
    console.error(e);
  }
})();
