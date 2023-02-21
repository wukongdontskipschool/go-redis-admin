package consts

const ENV_CONF string = "env/env"                                         // 本地配置路径
const ENV_CONF_PSALT string = "pSalt"                                     // 用户密码盐
const ENV_CONF_CRYPTOKEY string = "cryptoKey"                             // 加密盐 要求16个字符
const ENV_CONF_JWT_KEY string = "jetKey"                                  // jwt密钥
const ENV_CONF_GIN_MODE string = "ginMode"                                // ginmode
const JWT_CLAIMS string = "jwtClaims"                                     // gin ctx jwt验证结构体
const DB_RD_AD_CONF string = "databases/redisAdmin"                       // 数据库配置路径
const DB_RD_AD_CONF_TAG_AD string = "admin"                               // 数据库下标
const CASBIN_RBAC_CONF string = "internal/configs/casbin/rbacModels.conf" // rbac
const MENU_RD_LIST_ID uint = 99                                           // redis列表id
const ROLE_PRE string = "role_"                                           // rbac role前缀
const SUP_ADMIN_RID uint = 1                                              // 超级管理员角色id
const SUP_ADMIN_ID uint = 1                                               // 初始管理员id
