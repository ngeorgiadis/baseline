<template>
  <div class="main">
    <aside class="menu">
      <p class="menu-label">
        Toolbar
      </p>
      <ul class="menu-list">
        <li><a @click="fetchAndLoadServerData">Get server data</a></li>
        <!-- <li><a @click="saveGraph">Save to localstorage</a></li>
        <li><a @click="loadGraph">Load from localstorage</a></li> -->
      </ul>
      <p class="menu-label">
        Layout
      </p>
      <ul class="menu-list">
        <li>
          <a
            ><select class="select" v-model="layoutType">
              <option value="fruchterman">fruchterman</option>
              <option value="gForce">gForce</option>
              <option value="mds">mds</option>
              <option value="radial">radial</option>
              <option value="force">force</option>
            </select></a
          >
        </li>
        <li><a @click="fitToArea">Fit to view</a></li>
      </ul>
      <p class="menu-label">
        Subgraph
      </p>
      <ul class="menu-list">
        <li>
          <a
            ><select class="select" v-model="subgraph">
              <option v-for="i in 50" :key="i" :value="i">{{ i }}</option>
            </select></a
          >
        </li>
      </ul>
      <p class="menu-label">
        Options
      </p>
      <ul class="menu-list">
        <li>
          <a>
            <label class="checkbox"
              ><input class="checkbox" type="checkbox" v-model="getMaxCore" />
              Get Max K-Core subgraph
            </label>
          </a>
        </li>
      </ul>
      <p class="menu-label">
        Info
      </p>
      <div>
        <div>Max Dom Scrore: {{ stats.max_dom }}</div>
        <div>Min Dom Scrore: {{ stats.min_dom }}</div>

        <div>Max-Min: {{ stats.max_dom - stats.min_dom }}</div>

        <div>Avg Dom Scrore: {{ stats.avg_dom | toFixed }}</div>
        <div>Max k-Core: {{ stats.max_k_core }}</div>
        <p>
          R1: <b>{{ rank1 }}</b>
        </p>
      </div>
    </aside>
    <div class="graph" id="graph-container"></div>
    <div class="data-nodes table-container">
      <table class="table">
        <thead>
          <tr>
            <th>Label</th>
            <th>Domscore</th>
            <th>PC</th>
            <th>CN</th>
            <th>HI</th>
            <th>PI</th>
          </tr>
        </thead>
        <tr v-for="n in sortedNodes" :key="n.id">
          <td>{{ n.label }}</td>
          <td>{{ n.attrs.domscore }}</td>
          <td>{{ n.attrs.pc }}</td>
          <td>{{ n.attrs.cn }}</td>
          <td>{{ n.attrs.hi }}</td>
          <td>{{ n.attrs.pi }}</td>
        </tr>
      </table>
    </div>

    <!-- <button class="button" @click="fetchAndLoadServerData">
      Get from server
    </button>
    <button class="button" @click="saveGraph">Save to localstorage</button>
    <button class="button" @click="loadGraph">Load from localstorage</button>
    <select class="select" v-model="layoutType">
      <option value="fruchterman">fruchterman</option>
      <option value="gForce">gForce</option>
      <option value="mds">mds</option>
      <option value="radial">radial</option>
      <option value="force">force</option>
    </select>

    <select class="select" v-model="subgraph">
      <option v-for="i in 50" :key="i" :value="i">{{ i }}</option>
    </select>

    <input class="checkbox" type="checkbox" v-model="getMaxCore" /> -->
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
      subgraph: 1,
      getMaxCore: false,
      stats: {},
      graphData: {},
    };
  },

  async mounted() {
    this.createGraph({});
    let self = this;
    window.addEventListener("resize", function() {
      let elem = document.getElementById("graph-container");
      // let h = document.getElementById("graph-container").clientHeight;
      // let w = document.getElementById("graph-container").clientWidth;

      self.graph.changeSize(elem.clientWidth, elem.clientHeight);
    });
  },

  computed: {
    sortedNodes() {
      if (!this.graphData.nodes) {
        return [];
      }

      let nodes = this.graphData.nodes;

      nodes.sort((a, b) => {
        return +b.attrs.domscore - +a.attrs.domscore;
      });

      return nodes;
    },

    rank1() {
      let s = this.graphData.stats;
      if (!s) {
        return "";
      }

      let r = 0;

      if (s.max_k_core > 0) {
        r = s.max_k_core * s.avg_dom;
      } else {
        r = 2 * s.avg_dom;
      }

      return r.toFixed(3);
    },
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

      // this.createGraph(d);
    },

    async fetchServerData() {
      let url = this.getMaxCore
        ? `/app/graph/${this.subgraph - 1}/maxcore`
        : `/app/graph/${this.subgraph - 1}`;

      const response = await fetch(url);
      const remoteData = await response.json();
      this.stats = remoteData.stats;

      this.graphData = remoteData;

      let d = {
        nodes: remoteData.nodes.map((x) => {
          let clusterId = +x.cluster;

          // let c = x.shared
          //   ? "#aaa"
          //   : this.colors[clusterId % this.colors.length];

          let c = this.colors[clusterId % this.colors.length];

          let size = x.is_init ? 40 : 15;

          let label = `${x.label} [${x.attrs.domscore}, ${x.degree}, ${x.id}]`;

          return {
            ...x,
            label: label,
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

    fitToArea() {
      this.graph.fitView();
    },
  },

  watch: {
    layoutType: function(newVal) {
      this.graph.updateLayout({
        type: newVal,
        preventOverlap: true,
        gpuEnabled: true,
        workerEnabled: true,
      });
    },

    subgraph: function() {
      this.fetchAndLoadServerData();
    },

    getMaxCore: function() {
      this.fetchAndLoadServerData();
    },
  },

  filters: {
    toFixed(v) {
      if (!v) {
        return v;
      }

      return v.toFixed(2);
    },
  },
};
</script>

<style lang="scss">
.main {
  display: grid;
  grid-template-columns: 250px auto 550px;
}

.graph {
  display: grid;
  justify-content: center;
  height: calc(100vh - 25px);
}

.graph canvas {
  border: 1px solid #eee;
}
</style>
