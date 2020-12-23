<template>
  <div>
    <button @click="fetchAndLoadServerData">FETCH</button>
    <button @click="saveGraph">SAVE</button>
    <button @click="loadGraph">LOAD</button>
    <select v-model="layoutType">
      <option value="fruchterman">fruchterman</option>
      <option value="gForce">gForce</option>
      <option value="mds">mds</option>
      <option value="radial">radial</option>
      <option value="force">force</option>
    </select>

    <div class="graph" id="graph-container"></div>
  </div>
</template>

<script>
import G6 from "@antv/g6";
export default {
  data() {
    return {
      colors: [
        "#F44336",
        "#E91E63",
        "#9C27B0",
        "#2196F3",
        "#00BCD4",
        "#4CAF50",
        "#FFEB3B",
        "#CDDC39",
        "#FF9800",
        "#795548",
        "#607D8B",
        "#3F51B5",
        "#FFC107",
      ],
      layoutType: "fruchterman",
      graph: null,
    };
  },

  async mounted() {
    this.createGraph({});
    let self = this;
    window.addEventListener("resize", function () {
      let h = document.getElementById("graph-container").clientHeight;
      let w = document.getElementById("graph-container").clientWidth;

      self.graph.changeSize(w, h);
    });
  },

  methods: {
    createGraph(graphData) {
      let h = document.getElementById("graph-container").clientHeight;
      let w = document.getElementById("graph-container").clientWidth;

      this.graph = new G6.Graph({
        container: "graph-container",
        width: w,
        height: h,

        defaultNode: {
          size: 15,
          labelCfg: {
            style: {
              fontSize: 10,
              fontFamily: "Consolas",
            },
          },
        },

        modes: {
          default: [
            // {
            //   type: "drag-node",
            //   enableDelegate: true,
            //   shouldBegin: (e) => {
            //     // Do not allow the node with id 'node1' to be dragged
            //     // if (e.item && e.item.getModel().id === "node1") return false;
            //   },
            // },
            {
              type: "drag-node",
              enableDelegate: true,
            },
            "drag-canvas",
            "zoom-canvas",
            "brush-select",
            {
              type: "activate-relations",
              resetSelected: true,
              trigger: "click",
            },
          ],
        },

        layout: {
          type: this.layoutType,

          preventOverlap: true,
          maxIteration: 2000,
          gravity: 1,
          clustering: true,
          clusterGravity: 10,
          workerEnabled: true,
          gpuEnabled: true,
        },
        fitView: true,
        linkCenter: true,
      });

      this.graph.data(graphData); // Load the data defined in Step 2
      this.graph.render(); // Render the graph
    },

    async fetchAndLoadServerData() {
      let d = await this.fetchServerData();
      // this.graph.destroyLayout();
      this.graph.data(d);

      // this.graph.node((node) => {
      //   debugger;
      //   let d = 0;
      //   if (node) {
      //     d = this.graph.getNodeDegree(node.id);
      //   }

      //   return {
      //     ...node,
      //     label: `${node.label} (${d})`,
      //   };
      // });

      this.graph.render();
      this.graph.fitView();

      let n = this.graph.getNodes();
      console.log(n);
      debugger;
      // this.createGraph(d);
    },

    async fetchServerData() {
      const response = await fetch("/app/graph/1");
      const remoteData = await response.json();

      let d = {
        nodes: remoteData.nodes.map((x) => {
          let clusterId = +x.cluster;

          let c = x.shared
            ? "#aaa"
            : this.colors[clusterId % this.colors.length];

          let size = x.is_init ? 40 : 15;

          return {
            ...x,
            label: `${x.label} - ${x.attrs.domscore}`,
            size,
            style: {
              fill: c,
            },
          };
        }),
        edges: remoteData.edges,
      };

      return d;
    },

    saveGraph() {
      if (confirm("Overwrite??")) {
        let d = this.graph.save();
        localStorage.setItem("USER_GRAPH", JSON.stringify(d));
      }
    },

    loadGraph() {
      let s = localStorage.getItem("USER_GRAPH");
      let d = JSON.parse(s);
      this.graph.destroyLayout();
      this.graph.changeData(d);
      this.graph.fitView();
    },
  },

  watch: {
    layoutType: function (newVal) {
      this.graph.updateLayout({
        type: newVal,
        preventOverlap: true,
        gpuEnabled: true,
        workerEnabled: true,
      });
    },
  },
};
</script>

<style lang="scss">
.graph {
  display: grid;
  justify-content: center;
  height: calc(100vh - 25px);
}

.graph canvas {
  border: 1px solid #eee;
}
</style>