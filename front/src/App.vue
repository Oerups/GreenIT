<template>
  <div>
    <h1>Indice de fragilité numérique</h1>

    <div class="flex w-100 mt-3">
      <div class="left">
        <h2>Recherche</h2>

        <label for="iRegions">Régions</label>
        <input type="text" v-model="region" list="regions" id="iRegions">
        <datalist id="regions">
          <option v-for="(region, key) in regions" :key="key">{{ region }}</option>
        </datalist>

        <label for="iDepartments" class="mt-1">Départements</label>
        <input type="text" v-model="department" list="departments" id="iDepartments">
        <datalist id="departments">
          <option v-for="(department, key) in departments" :key="key">{{ department }}</option>
        </datalist>

        <label for="iMunicipalities" class="mt-1">Communes</label>
        <input type="text" v-model="municipality" list="municipalities" id="iMunicipalities">
        <datalist id="municipalities">
          <option v-for="(value, key) in municipalities" :key="key">{{ value }}</option>
        </datalist>

        <button class="mt-1" @click="search">Chercher</button>
      </div>

      <div class="right">
        <div style="width: 100%" v-if="result.length">
          <h2>Résultat</h2>


          <table id="tableData">
            <tr>
              <th>Commune</th>
              <th>Code iris</th>
              <th>Nom iris</th>
              <th>Département</th>
              <th>Région</th>
              <th>Population</th>
              <th>SCORE GLOBAL</th>
              <th>ACCES A L'INFORMATION</th>
              <th>ACCÈS AUX INTERFACES NUMERIQUES</th>
              <th>COMPETENCES ADMINISTATIVES</th>
              <th>COMPÉTENCES NUMÉRIQUES / SC..</th>
              <th>GLOBAL ACCES</th>
              <th>GLOBAL COMPETENCES</th>
            </tr>
            <tr v-for="(value, key) in result" :key="key">
              <td>{{ value["libcom"] }}</td>
              <td>{{ value["code_iris"] }}</td>
              <td>{{ value["nom_iris"] }}</td>
              <td>{{ value["libdep"] }}</td>
              <td>{{ value["libreg"] }}</td>
              <td>{{ value["pop"] }}</td>
              <td>{{ value["score_global"] }}</td>
              <td>{{ value["acces_infos"] }}</td>
              <td>{{ value["acces_interf"] }}</td>
              <td>{{ value["competences_admin"] }}</td>
              <td>{{ value["competences_num"] }}</td>
              <td>{{ value["global_acces"] }}</td>
              <td>{{ value["global_comp"] }}</td>
            </tr>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      region: '',
      regions: [],
      department: '',
      departments: [],
      municipality: '',
      municipalities: [],
      result: [],
    }
  },
  methods: {
    search() {
      const values = {};
      if (this.region != "") {
        values.regions = this.region;
      }
      if (this.municipality != "") {
        values.municipalities = this.municipality;
      }
      if (this.department != "") {
        values.departments = this.department;
      }

      fetch(`http://localhost:3000/search?` + new URLSearchParams(values), {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        },
        method: "GET",
      })
          .then((res) => {
            if (res.ok) {
              return res.json();
            }
          })
          .then((data) => {
            console.log(data);
            this.result = data;
          })
          .catch(() => {
            console.log("Error");
          })
    }
  },
  beforeMount() {
    fetch(`http://localhost:3000/data`, {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      method: "GET"
    })
        .then((res) => {
          if (res.ok) {
            return res.json()
          }
        })
        .then((data) => {
          this.regions = data.regions;
          this.departments = data.departments;
          this.municipalities = data.municipalities;
        })
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}

h1 {
  text-align: center;
}

input {
  width: fit-content;
}

.flex {
  display: flex;
}

.mt-1 {
  margin-top: 25px;
}

.mt-3 {
  margin-top: 75px;
}

.left {
  width: 20%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.right {
  width: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;
  overflow: scroll;
}

.w-100 {
  width: 100%;
}

 #tableData {
   font-family: Arial, Helvetica, sans-serif;
   border-collapse: collapse;
   width: 100%;
 }

#tableData td, #tableData th {
  border: 1px solid #ddd;
  padding: 8px;
}

#tableData tr:nth-child(even){background-color: #f2f2f2;}

#tableData tr:hover {background-color: #ddd;}

#tableData th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #04AA6D;
  color: white;
}
</style>
