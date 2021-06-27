<template>
  <div class="app">
    <!-- 头部 -->
    <div class="header">
      <span class="header_left"><i class="el-icon-school" /> 教育部全国高校名单（截止2020-06-30）</span>
      <div class="header_right">
        <el-tabs v-model="activeName" @tab-click="handleClick">
          <el-tab-pane label="按高校查询" name="first" />
          <el-tab-pane label="按地区查询" name="second" />
        </el-tabs>
      </div>
    </div>

    <!-- 中部 -->
    <el-container class="container">
      <!-- 导航栏 -->
      <el-aside class="container_left" width="200px">
        <el-card> 
          <el-menu :default-openeds="['1']" class="container_left_menu">

            <el-submenu index="1">
              <template #title><span class="title">+&nbsp;&nbsp;高校类型</span></template>
              <el-menu-item-group v-for="(type_item, type_key) in school_type" :key="type_key" v-model="submenu.school_type" class="submenu">
                <el-menu-item :index="'1-' + type_key.toString()">{{ type_item.value }}</el-menu-item>
              </el-menu-item-group>
            </el-submenu>

            <el-submenu index="2">
              <template #title><span class="title">+&nbsp;&nbsp;所在地区</span></template>
              <el-menu-item-group v-for="(province_item, province_key) in school_province" :key="province_key" v-model="submenu.school_province" class="submenu">
                <el-menu-item :index="'2-' + province_key.toString()">{{ province_item.value }}</el-menu-item>
              </el-menu-item-group>
            </el-submenu>

          </el-menu>
        </el-card>
      </el-aside>

    

      <!-- 表格栏 -->
      <el-card class="container_right">
        <el-header><span class="dot">●</span><span>一流大学</span></el-header>
        <el-main>
          <el-table :data="tableData" stripe border show-overflow-tooltip :cell-style="{ fontSize: '13px' }" style="z-index: 10;">
            <el-table-column prop="school_id" label="№" width="30" align="center" />
            <el-table-column prop="school_code" label="学校标识码" width="90" align="center" />
            <el-table-column prop="school_name" label="学校名称" width="200" align="center" />
            <el-table-column prop="school_province" label="所在地区" width="130" align="center" />
            <el-table-column prop="school_city" label="所在城市" width="90" align="center" />
            <el-table-column prop="school_department" label="主管部门" min-width="90" align="center" />
            <el-table-column prop="school_level" label="办学层次" width="80" align="center" />
            <el-table-column prop="school_type" label="高校类型" align="center" />
          </el-table>
        </el-main>
      </el-card>
    </el-container>
  </div>
</template>

<script>

export default {
  data() {
    const item = {
      school_id: 1,
      school_code: 4111010001,
      school_name: '中国人民解放军国防科技大学',
      school_province: '新疆维吾尔自治区',
      school_city: '乌鲁木齐',
      school_department: '中国共产党中央军事委员会',
      school_level: '本科',
      school_type: '一流大学A类,985,211'
    }
    
    return {
      school_province: [
        { key: '1', value: '北京市' },
        { key: '2', value: '广东省' },
        { key: '3', value: '山东省' },
        { key: '4', value: '江苏省' },
        { key: '5', value: '河南省' },
        { key: '6', value: '上海市' },
        { key: '7', value: '河北省' },
        { key: '8', value: '浙江省' },
        { key: '9', value: '天津市' },
        { key: '10', value: '四川省' },
        { key: '11', value: '陕西省' },
        { key: '12', value: '湖南省' },
        { key: '13', value: '重庆市' },
        { key: '14', value: '吉林省' },
        { key: '15', value: '云南省' },
        { key: '16', value: '福建省' },
        { key: '17', value: '安徽省' },
        { key: '18', value: '海南省' },
        { key: '19', value: '江西省' },
        { key: '20', value: '湖北省' },
        { key: '21', value: '辽宁省' },
        { key: '22', value: '山西省' },
        { key: '23', value: '台湾省' },
        { key: '24', value: '黑龙江' },
        { key: '25', value: '西藏区' },
        { key: '26', value: '青海省' },
        { key: '27', value: '贵州省' },
        { key: '28', value: '甘肃省' },
        { key: '29', value: '宁夏回族自治区' },
        { key: '30', value: '广西壮族自治区' },
        { key: '31', value: '新疆维吾尔自治区' },
        { key: '32', value: '内蒙古自治区' },
        { key: '33', value: '澳门特别行政区' },
        { key: '34', value: '香港特别行政区' }
      ],
      school_type: [
        { key: '1', value: '一流大学' },
        { key: '2', value: '一流学科' },
        { key: '3', value: '985高校' },
        { key: '4', value: '211高校' }
      ],
      submenu: {
        school_type: '',
        school_province: ''
      },
      activeName: 'second',
      tableData: Array(50).fill(item),
      school_id: undefined,
      school_code: undefined,
      school_name: undefined,
      // school_province: undefined,
      school_city: undefined,
      school_department: undefined,
      school_level: undefined,
      // school_type: undefined,
      school_website: undefined
    }
  },

  methods: {
    handleClick(tab, event) {
      console.log(tab, event)
    }
  }
}
</script>

<style lang='scss'>
// 全局样式
  .app {
    width: 1150px;
    margin: auto;
  }

// 头部样式
  .header {
    position: relative;
    height: 50px;
    background-color: #f2f8fe;
    border-radius: 6px;

    .el-icon-school {
      margin-left: 10px;
    }

    .header_left {
      position: absolute;
      margin-top: 13px;

      .el-icon-school {
        color: #4285f4;
      }
    }

    .header_right {
      float:right;
      width: 200px;
    }
  }

// 中部样式
  .container {
    margin-top: 15px;
    // height: 500px;

    .container_left {

      .container_left_menu {
        border-right: 0px;

        .submenu {
          background-color: #f8f8f8;
        }
      }
    }

    .container_right {
      margin-left: 15px;
      width: 100%;

      .dot {
        color: #4285f4;
        margin-right: 10px;
      }

      .el-header {
        border-bottom: 2px solid #4285f4;
        color: #333;
        line-height: 60px;
      }
    }
  }

// element-ui卡片单独样式
  .el-card__body {
    padding: 0;
  }
  .el-main {
    padding: 0;
  }
  .el-table .cell {
    padding: 2px;
  }
  .cell {
    padding: 2px;
  }
  .el-table--border td:first-child .cell, .el-table--border th:first-child .cell {
    padding-left: 2px;
}
</style>
