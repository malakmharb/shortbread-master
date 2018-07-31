import axios from "axios";

const instance = axios.create();
instance.defaults.headers.common['X-CSRF-Token'] = $('meta[name=csrf-token]').attr('content');

module.exports = instance;

