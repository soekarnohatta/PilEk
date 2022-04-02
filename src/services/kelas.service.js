import {instance} from "@/utils/axios";

class KelasService {
    getKelas() {
        return instance.get('/kelas/get');
    }

    getAll() {
        return instance.get('/kelas/getall',);
    }
}

export default new KelasService();