<template>
<div>
    <h1>Indice de fragilité numérique</h1>

    <div class="flex w-100 mt-3">
        <div class="left">
            <h2>Recherche</h2>

            <label for="iMunicipalities">Communes</label>
            <input type="text" v-model="municipality" list="municipalities" id="iMunicipalities">
            <datalist id="municipalities">
                <option v-for="(value, key) in municipalities" :key="key">{{ value }}</option>
            </datalist>

            <label for="iRegions" class="mt-1">Régions</label>
            <input type="text" v-model="region" list="regions" id="iRegions">
            <datalist id="regions">
                <option v-for="(region, key) in regions" :key="key">{{ region }}</option>
            </datalist>

            <label for="iDepartments" class="mt-1">Départements</label>
            <input type="text" v-model="department" list="departments" id="iDepartments">
            <datalist id="departments">
                <option v-for="(department, key) in departments" :key="key">{{ department }}</option>
            </datalist>

            <button class="mt-1" @click="search">Chercher</button>
        </div>

        <div class="right">
            <div v-if="result.length">
                <h2>Résultat</h2>

                <div style="margin-top: 10px;" v-for="(value, key) in result" :key="key">
                    <div>Commune : {{ value["Libcom"] }}</div>
                    <div>Score accès à l'information : </div>
                    <div>Score accès aux interfaces numériques : </div>
                    <div>Score compétences administratives : </div>
                    <div>Score compétences numériques : </div>
                    <div>Score global : </div>
                </div>
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
            if(this.region != "") {
                values.regions = this.region;
            }
            if(this.municipality != "") {
                values.municipalities = this.municipality;
            }
            if(this.department != "") {
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
}

.w-100 {
    width: 100%;
}
</style>
