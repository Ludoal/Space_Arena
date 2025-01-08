
local M = {}

-- controller
M.SHOW_MENU = hash("show_menu")
M.SHOW_ARENA = hash("show_arena")
M.PROXY_LOADED = hash("proxy_loaded")
M.PROXY_UNLOADED = hash("proxy_unloaded")
M.QUIT = hash("quit")

-- game
M.GAME_TIME = hash("game_time")
M.GAME_OVER = hash("game_over")
M.TIME_UP = hash("time_up")

-- display
M.UPDATE_LAYERS = hash("update_layers")
M.UPDATE_FPS = hash("update_fps")
M.UPDATE_SHIP_DAMPER = hash("update_ship_damper")
M.UPDATE_SHIP_STATS = hash("update_ship_stats")
M.UPDATE_SHIP_INPUT_1 = hash("update_ship_input_1")
M.UPDATE_SHIP_INPUT_2 = hash("update_ship_input_2")

-- camera
M.UPDATE_CAMERA = hash("update_camera") -- duplicate

-- physics
M.CONTACT_POINT_RESPONSE = hash("contact_point_response")
M.COLLISION_RESPONSE = hash("collision_response")
M.TRIGGER_RESPONSE = hash("trigger_response")

-- player
M.UPDATE_SCORE = hash("update_score")
M.UPDATE_HEALTH = hash("update_health")
M.UPDATE_AMMO = hash("update_ammo")
M.DEATH = hash("detah")

-- powerups
M.ADD_AMMO = hash("add_ammo")
M.ADD_LIFE = hash("add_life")
M.ADD_GUN = hash("add_gun")

-- bullet
M.BULLET_HIT = hash("bullet_hit")

-- border
M.FIELD_FORCE = hash("field_force")

-- enemy
M.ENEMY_TARGET = hash("enemy_target")

-- input
M.THR_F = hash("KEY_W")
M.THR_B = hash("KEY_S")
M.THR_L = hash("KEY_Q")
M.THR_R = hash("KEY_E")
M.ROT_L = hash("KEY_A")
M.ROT_R = hash("KEY_D")
M.DAMPER = hash("KEY_Z")
M.FIRE = hash("KEY_SPACE")
M.ZOOM_IN = hash("KEY_UP")
M.ZOOM_OUT = hash("KEY_DOWN")
M.ESCAPE = hash("KEY_ESC")
M.CLICK_L = hash("CLICK_L")
M.CLICK_R = hash("CLICK_R")

return M
