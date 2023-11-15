<template>
    <div>
        <Menu />
        <div class="container">
            <h1 class="m0">Article</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage your Article</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-Article">
                        <button class="bt-green">+ Add</button>
                    </router-link>
                    <td></td>
                </div>
            </div>
            <td>
                <div class="container">
                    <div class="row">
                        <div class="col center">
                            <router-link to="/list-user">
                                <div>
                                    List Product
                                    <hr class="uline-green">
                                </div>
                            </router-link>
                        </div>
                        <div class="col center">
                            <router-link to="/cart">
                                <div>
                                    Voucher
                                    <hr class="uline-grey">
                                </div>
                            </router-link>
                        </div>
                        <div class="col center">
                            <router-link to="/cart">
                                <div>
                                    Banner
                                    <hr class="uline-grey">
                                </div>
                            </router-link>
                        </div>
                        <div class="col center">
                            <router-link to="/cart">
                                <div>
                                    Plannogram
                                    <hr class="uline-grey">
                                </div>
                            </router-link>
                        </div>
                        <div class="col center">
                            <router-link to="/list-Article">
                                <div>
                                    Article
                                    <hr class="uline-grey">
                                </div>
                            </router-link>
                        </div>
                        <div class="col center">
                            <router-link to="/cart">
                                <div>
                                    Notification
                                    <hr class="uline-grey">
                                </div>
                            </router-link>
                        </div>
                        <div class="col center">
                            <router-link to="/cart">
                                <div>
                                    Delivery Method
                                    <hr class="uline-grey">
                                </div>
                            </router-link>
                        </div>
                    </div>
                </div>
            </td>
            <div class="row bg-blue mt10 p10">
                <div class="col">
                    <input class="search-field" type="text" id="deskripsi"
                        placeholder="Search by invoice number, name, amount ..." />
                </div>
                <div class="col flex">
                    <button class="bt-blue">Filter</button>
                </div>
            </div>
            <div class="mt50">
                <table>
                    <thead>
                        <tr>
                            <th>ARTICLE NAME</th>
                            <th>BROADCAST SCHEDULE</th>
                            <th>VALID FROM</th>
                            <th>VALID UNTIL</th>
                            <th>ARTICLE STATUS</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item of list_article" :key="item.id">
                            <td class="center">{{ item.article }}</td>
                            <td class="center">{{ item.bs }}</td>
                            <td class="center">{{ item.vf }}</td>
                            <td class="center">{{ item.vu }}</td>                    
                            <td class="center">
                                <div v-if="item.as === 1">
                                    <label class="label-green">Active</label>
                                </div>
                                <div v-if="item.as === 2">
                                    <label class="label-green">Active</label>
                                </div>
                                <!-- <div v-if="item.role === 3">
                                    <label class="label-yellow">Marketing</label>
                                </div> -->
                            </td>
                            <td class="center">
                                <div class="mt5">
                                    <router-link to="/edit-Article">
                                        <img src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link to="/detail-Article">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img src="../../assets/icons/delete.svg">
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script>
import Menu from "../../components/menu.vue"

export default {
    name: 'list-Article',
    components: {
        Menu
    },
    data() {
        return {
            
            list_article: [
                {
                    id: "1",
                    article: "Vioni Wita Elya",
                    bs: "viowelya",
                    vf: "vioniwitaelya@gmail.com",
                    vu : "04-06-2021",
                    as: 1
                },
                {
                    id: "2",
                    article: "Haifah Nur Zakiah",
                    bs: "viowelya",
                    vf: "04-06-2021",
                    vu : "04-06-2021",
                    as: 2
                },
                // {
                //     id: "3",
                //     name: "Nabila Hasyimie",
                //     username: "bilamie",
                //     email: "bilamie07@gmail.com",
                //     role: 3
                // },
                // {
                //     id: "4",
                //     name: "Naufal Al-Ghifary",
                //     username: "naufary",
                //     email: "naufary@gmail.com",
                //     role: 3
                // },
                // {
                //     id: "5",
                //     name: "Hafidz Zaki Makarim",
                //     username: "fidzki",
                //     email: "fidzki@gmail.com",
                //     role: 3
                // }
            ]
        }
    }
}
</script>