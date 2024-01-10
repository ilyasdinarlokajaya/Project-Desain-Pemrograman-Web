import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const state = {
  selected_menu: "user",
  user: {
    detail: {
      id: "",
      name: "",
      username: "",
      email: "",
      role: 0,
      image: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "banner",
  banner: {
    detail: {
      id: "",
      bannerName: "",
      broadcastDay: "",
      validFrom: "",
      validUntil: "",
      image: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "notification",
  notification: {
    detail: {
      id: "",
      title: "",
      konten: "",
      broadcastschedule: "",
      notificationstatus: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "product",
  product: {
    detail: {
      id: "",
      product_code: "",
      product_name: "",
      purchase: 0,
      productdescription: "",
      unit: 0,
      minpurchase: 0,
      maxpurchase: 0,
      category: "",
      image: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "plannogram",
  plannogram: {
    detail: {
      id: "",
      plannogramName: "",
      broadcastDay: null,
      validFrom: "",
      validUntil: "",
      listproduct: null,
      image: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "voucher",
  voucher: {
    detail: {
      id: "",
      vouchercode: "",
      vouchername: "",
      description: "",
      type: "",
      mintransaction: "",
      amount: "",
      validfrom: "",
      validuntil: "",
      Image: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "article",
  article: {
    detail: {
      id: "",
      judul_artikel: "",
      konten: "",
      jadwal_broadcast: "",
      valid_from: "",
      valid_until: "",
      Image: ""
    },
    list: [],
    search: ""
  },
  selected_menu: "delivery-method",
  delivery_method: {
    detail: {
      id: "",
      delivery_name: "",
      delivery_fee: "",
      estimation: "",
      Image: ""
    },
    list: [],
    search: ""
  }
};

const mutations = {
  setSelectedMenu(state, payload) {
    state.selected_menu = payload;
  },
  setDetailUser(state, payload) {
    state.user.detail = payload;
  },
  setListUser(state, payload) {
    state.user.list = payload;
  },
  setSearchBanner(state, payload) {
    console.log("set asjdkhaskjld");
    state.banner.search = payload;
  },
  setDetailBanner(state, payload) {
    state.banner.detail = payload;
  },
  setListBanner(state, payload) {
    state.banner.list = payload;
  },
  setSearchNotification(state, payload) {
    console.log("set asjdkhaskjld");
    state.notification.search = payload;
  },
  setDetailNotification(state, payload) {
    state.notification.detail = payload;
  },
  setListNotification(state, payload) {
    state.notification.list = payload;
  },
  setSearchProduct(state, payload) {
    console.log("set asjdkhaskjld");
    state.product.search = payload;
  },
  setDetailProduct(state, payload) {
    state.product.detail = payload;
  },
  setListProduct(state, payload) {
    state.product.list = payload;
  },
  setSearchPlannogram(state, payload) {
    console.log("set asjdkhaskjld");
    state.plannogram.search = payload;
  },
  setDetailPlannogram(state, payload) {
    state.plannogram.detail = payload;
  },
  setListPlannogram(state, payload) {
    state.plannogram.list = payload;
  },
  setDetailVoucher(state, payload) {
    state.voucher.detail = payload;
  },
  setListVoucher(state, payload) {
    state.voucher.list = payload;
  },
  setSearchVoucher(state, payload) {
    console.log("set asjdkhaskjld");
    state.voucher.search = payload;
  },
  setDetailArticle(state, payload) {
    state.article.detail = payload;
  },
  setListArticle(state, payload) {
    state.article.list = payload;
  },
  setSearchArticle(state, payload) {
    console.log("set asjdkhaskjld");
    state.article.search = payload;
  },
  setDetailDeliveryMethod(state, payload) {
    state.delivery_method.detail = payload;
  },
  setListDeliveryMethod(state, payload) {
    state.delivery_method.list = payload;
  },
  setSearchDeliveryMethod(state, payload) {
    console.log("set asjdkhaskjld");
    state.delivery_method.search = payload;
  }
};

const actions = {
  createUser: ({ commit, state, dispatch }, payload) => {
    try {
      axios.post("http://localhost:3001/user", payload);
    } catch (error) {}
  },
  addUser: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/user", payload);
      dispatch("listUser");
    } catch (e) {
      console.log(e);
    }
  },
  listUser: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/user");
          if (state.user.search) {
            let filter = [
              res.data.find(user => user.username === state.user.search)
            ];
            commit("setListUser", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListUser", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailUser: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/user/" + id);
          commit("setDetailUser", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteUser: ({ commit, state, dispatch }, id) => {
    try {
      const delUser = async () => {
        try {
          const res = await axios.delete("http://localhost:3001/user/" + id);
          if (res.status === 200) {
            dispatch("listUser");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delUser();
    } catch (e) {
      console.log(e);
    }
  },
  editUser: ({ commit, state, dispatch }, data) => {
    try {
      const delUser = async () => {
        try {
          const res = await axios.put("http://localhost:3001/user/" + data.id, {
            name: data.name,
            username: data.username,
            email: data.email,
            role: parseInt(data.role),
            image: data.image
          });
          if (res.status === 200) {
            dispatch("listUser");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delUser();
    } catch (e) {
      console.log(e);
    }
  },
  createBanner: ({ commit, state, dispatch }, payload) => {
    try {
      axios.post("http://localhost:3001/banner", payload);
    } catch (error) {}
  },
  addBanner: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/banner", payload);
      dispatch("listBanner");
    } catch (e) {
      console.log(e);
    }
  },
  listBanner: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/banner");
          if (state.banner.search) {
            let filter = [
              res.data.find(banner => banner.bannerName === state.banner.search)
            ];
            commit("setListBanner", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListBanner", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailBanner: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/banner/" + id);
          commit("setDetailBanner", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteBanner: ({ commit, state, dispatch }, id) => {
    try {
      const delBanner = async () => {
        try {
          const res = await axios.delete("http://localhost:3001/banner/" + id);
          if (res.status === 200) {
            dispatch("listBanner");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delBanner();
    } catch (e) {
      console.log(e);
    }
  },
  editBanner: ({ commit, state, dispatch }, data) => {
    try {
      const delBanner = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/banner/" + data.id,
            {
              bannerName: data.bannerName,
              broadcastDay: data.broadcastDay,
              validFrom: data.validFrom,
              validUntil: data.validUntil,
              image: data.image
            }
          );
          if (res.status === 200) {
            dispatch("listBanner");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delBanner();
    } catch (e) {
      console.log(e);
    }
  },
  createNotification: ({ commit, state, dispatch }, payload) => {
    try {
      axios.post("http://localhost:3001/notification", payload);
    } catch (error) {}
  },
  addNotification: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/notification", payload);
      dispatch("listNotification");
    } catch (e) {
      console.log(e);
    }
  },
  listNotification: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/notification");
          if (state.notification.search) {
            let filter = [
              res.data.find(
                notification => notification.title === state.notification.search
              )
            ];
            commit(
              "setListNotification",
              filter[0] !== undefined ? filter : []
            );
          } else {
            commit("setListNotification", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailNotification: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get(
            "http://localhost:3001/notification/" + id
          );
          commit("setDetailNotification", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteNotification: ({ commit, state, dispatch }, id) => {
    try {
      const delNotif = async () => {
        try {
          const res = await axios.delete(
            "http://localhost:3001/notification/" + id
          );
          if (res.status === 200) {
            dispatch("listNotification");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delNotif();
    } catch (e) {
      console.log(e);
    }
  },
  editNotification: ({ commit, state, dispatch }, data) => {
    try {
      const delNotification = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/notification/" + data.id,
            {
              title: data.title,
              konten: data.konten,
              broadcast_day: data.broadcast_day
            }
          );
          if (res.status === 200) {
            dispatch("listNotification");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delNotification();
    } catch (e) {
      console.log(e);
    }
  },
  createProduct: ({ commit, state, dispatch }, payload) => {
    try {
      axios.post("http://localhost:3001/product", payload);
    } catch (error) {}
  },
  addProduct: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/product", payload);
      dispatch("listProduct");
    } catch (e) {
      console.log(e);
    }
  },
  listProduct: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/product");
          if (state.product.search) {
            let filter = [
              res.data.find(
                product => product.product_code === state.product.search
              )
            ];
            commit("setListProduct", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListProduct", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailProduct: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/product/" + id);
          commit("setDetailProduct", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteProduct: ({ commit, state, dispatch }, id) => {
    try {
      const delProduct = async () => {
        try {
          const res = await axios.delete("http://localhost:3001/product/" + id);
          if (res.status === 200) {
            dispatch("listProduct");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delProduct();
    } catch (e) {
      console.log(e);
    }
  },
  editProduct: ({ commit, state, dispatch }, data) => {
    try {
      const delProduct = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/product/" + data.id,
            {
              product_code: data.product_code,
              product_name: data.product_name,
              purchase: parseInt(data.purchase),
              productdescription: data.productdescription,
              unit: parseInt(data.unit),
              minpurchase: parseInt(data.minpurchase),
              maxpurchase: parseInt(data.maxpurchase),
              category: data.category,
              image: data.image
            }
          );
          if (res.status === 200) {
            dispatch("listProduct");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delProduct();
    } catch (e) {
      console.log(e);
    }
  },

  addPlannogram: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/plannogram", payload);
      dispatch("listPlannogram");
    } catch (e) {
      console.log(e);
    }
  },
  listPlannogram: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/plannogram");
          if (state.plannogram.search) {
            let filter = [
              res.data.find(
                plannogram =>
                  plannogram.plannogramName === state.plannogram.search
              )
            ];
            commit("setListPlannogram", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListPlannogram", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailPlannogram: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/plannogram/" + id);
          commit("setDetailPlannogram", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deletePlannogram: ({ commit, state, dispatch }, id) => {
    try {
      const deletePlannogram = async () => {
        try {
          const res = await axios.delete(
            "http://localhost:3001/plannogram/" + id
          );
          if (res.status === 200) {
            dispatch("listPlannogram");
          }
        } catch (err) {
          console.error(err);
        }
      };
      deletePlannogram();
    } catch (e) {
      console.log(e);
    }
  },
  editPlannogram: ({ commit, state, dispatch }, data) => {
    try {
      const delPlannogram = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/plannogram/" + data.id,
            {
              plannogramName: data.plannogramName,
              broadcastDay: data.broadcastDay,
              validFrom: data.validFrom,
              validUntil: data.validUntil,
              listproduct: data.listproduct,
              image: data.image
            }
          );
          if (res.status === 200) {
            dispatch("listPlannogram");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delPlannogram();
    } catch (e) {
      console.log(e);
    }
  },
  createVoucher: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/voucher", payload);
      dispatch("listVoucher");
    } catch (error) {
      console.log(error);
    }
  },
  addVoucher: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/voucher", payload);
      dispatch("listVoucher");
    } catch (e) {
      console.log(e);
    }
  },
  listVoucher: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/voucher");
          if (state.voucher.search) {
            let filter = [
              res.data.find(
                voucher => voucher.vouchername === state.voucher.search
              )
            ];
            commit("setListVoucher", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListVoucher", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailVoucher: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/voucher/" + id);
          commit("setDetailVoucher", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteVoucher: ({ commit, state, dispatch }, id) => {
    try {
      const delVoucher = async () => {
        try {
          const res = await axios.delete("http://localhost:3001/voucher/" + id);
          if (res.status === 200) {
            dispatch("listVoucher");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delVoucher();
    } catch (e) {
      console.log(e);
    }
  },
  editVoucher: ({ commit, state, dispatch }, data) => {
    try {
      const delVoucher = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/voucher/" + data.id,
            {
              vouchercode: data.vouchercode,
              vouchername: data.vouchername,
              description: data.description,
              type: data.type,
              mintransaction: data.mintransaction,
              amount: data.amount,
              validfrom: data.validfrom,
              validuntil: data.validuntil,
              image: data.image
            }
          );
          if (res.status === 200) {
            dispatch("listVoucher");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delVoucher();
    } catch (e) {
      console.log(e);
    }
  },
  addArticle: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/article", payload);
      dispatch("listArticle");
    } catch (e) {
      console.log(e);
    }
  },
  listArticle: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/article");
          if (state.article.search) {
            let filter = [
              res.data.find(
                article => article.judul_artikel === state.article.search
              )
            ];
            commit("setListArticle", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListArticle", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailArticle: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/article/" + id);
          commit("setDetailArticle", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteArticle: ({ commit, state, dispatch }, id) => {
    try {
      const delArticle = async () => {
        try {
          const res = await axios.delete("http://localhost:3001/article/" + id);
          if (res.status === 200) {
            dispatch("listArticle");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delArticle();
    } catch (e) {
      console.log(e);
    }
  },
  editArticle: ({ commit, state, dispatch }, data) => {
    try {
      const delArticle = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/article/" + data.id,
            {
              judul_artikel: data.judul_artikel,
              jadwal_broadcast: data.jadwal_broadcast,
              konten: data.konten,
              valid_from: data.valid_from,
              valid_until: data.valid_until,
              image: data.image
            }
          );
          if (res.status === 200) {
            dispatch("listArticle");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delArticle();
    } catch (e) {
      console.log(e);
    }
  },
  addDeliveryMethod: ({ commit, state, dispatch }, payload) => {
    try {
      console.log("check payload", payload);
      axios.post("http://localhost:3001/delivery-method", payload);
      dispatch("listDeliveryMethod");
    } catch (e) {
      console.log(e);
    }
  },
  listDeliveryMethod: ({ commit, state }) => {
    try {
      const getData = async () => {
        try {
          const res = await axios.get("http://localhost:3001/delivery-method");
          if (state.delivery_method.search) {
            let filter = [
              res.data.find(
                delivery_method => delivery_method.delivery_name === state.delivery_method.search
              )
            ];
            commit("setListDeliveryMethod", filter[0] !== undefined ? filter : []);
          } else {
            commit("setListDeliveryMethod", res.data);
          }
        } catch (err) {
          console.error(err);
        }
      };
      getData();
    } catch (e) {
      console.log(e);
    }
  },
  detailDeliveryMethod: ({ commit, state }, id) => {
    try {
      const getDetail = async () => {
        try {
          const res = await axios.get("http://localhost:3001/delivery-method/" + id);
          commit("setDetailDeliveryMethod", res.data);
        } catch (err) {
          console.error(err);
        }
      };
      getDetail();
    } catch (e) {
      console.log(e);
    }
  },
  deleteDeliveryMethod: ({ commit, state, dispatch }, id) => {
    try {
      const delDeliveryMethod = async () => {
        try {
          const res = await axios.delete("http://localhost:3001/delivery-method/" + id);
          if (res.status === 200) {
            dispatch("listDeliveryMethod");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delDeliveryMethod();
    } catch (e) {
      console.log(e);
    }
  },
  editDeliveryMethod: ({ commit, state, dispatch }, data) => {
    try {
      const delDeliveryMethod = async () => {
        try {
          const res = await axios.put(
            "http://localhost:3001/delivery-method/" + data.id,
            {
              delivery_name: data.delivery_name,
              delivery_fee: data.delivery_fee,
              estimation: data.estimation,
              image: data.image
            }
          );
          if (res.status === 200) {
            dispatch("listDeliveryMethod");
          }
        } catch (err) {
          console.error(err);
        }
      };
      delDeliveryMethod();
    } catch (e) {
      console.log(e);
    }
  }
};

export default new Vuex.Store({
	state,
	mutations,
	actions
})
