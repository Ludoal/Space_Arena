
local M = {}

-- API
M.COLLISION_RESPONSE = hash("collision_response")
M.CONTACT_POINT_RESPONSE = hash("contact_point_response")
M.TRIGGER_RESPONSE = hash("trigger_response")

-- menu_gui
M.CLICK_L = hash("CLICK_L") -- input -- and target (if debug)

-- main_controller
M.SHOW_MENU = hash("show_menu")
M.SHOW_ARENA = hash("show_arena")
M.PROXY_LOADED = hash("proxy_loaded")
M.PROXY_UNLOADED = hash("proxy_unloaded")
M.QUIT = hash("quit")
M.ESCAPE = hash("KEY_ESC") -- input

-- arena_controller
M.UPDATE_LAYERS = hash("update_layers")
M.SPAWN_POWERUP = hash("spawn_powerup")
M.UPDATE_SCORE = hash("update_score") -- and arena_gui
M.HIGH_SCORE = hash("high_score") -- and arena_gui
M.GAME_OVER = hash("game_over") -- and arena_gui

-- arena_gui
M.GAME_TIME = hash("game_time")
M.TIME_UP = hash("time_up")
M.DEATH = hash("death")
M.UPDATE_FPS = hash("update_fps")
M.UPDATE_HEALTH = hash("update_health")
M.UPDATE_AMMO = hash("update_ammo")
M.UPDATE_SHIP_DAMPER = hash("update_ship_damper")
M.UPDATE_SHIP_STATS = hash("update_ship_stats")
M.UPDATE_SHIP_INPUT_1 = hash("update_ship_input_1")
M.UPDATE_SHIP_INPUT_2 = hash("update_ship_input_2")

-- asteroid
M.BULLET_HIT = hash("bullet_hit") -- and enemy, ship
M.BULLET_HIT_BY_SHIP = hash("bullet_hit_by_ship")

-- camera
M.UPDATE_CAMERA = hash("update_camera")
M.ZOOM_IN = hash("KEY_UP") -- input
M.ZOOM_OUT = hash("KEY_DOWN") -- input

-- enemy
M.ENEMY_TARGET = hash("enemy_target")

-- ship
M.REMOVE_LIFE = hash("remove_life")
M.ADD_LIFE = hash("add_life")
M.ADD_AMMO = hash("add_ammo")
M.ADD_GUN = hash("add_gun")
M.FIELD_FORCE = hash("field_force")
M.THR_F = hash("KEY_W") -- input
M.THR_B = hash("KEY_S") -- input
M.THR_L = hash("KEY_Q") -- input
M.THR_R = hash("KEY_E") -- input
M.ROT_L = hash("KEY_A") -- input
M.ROT_R = hash("KEY_D") -- input
M.DAMPER = hash("KEY_Z") -- input
M.FIRE = hash("KEY_SPACE") -- input

-- target
M.CLICK_R = hash("CLICK_R") -- (if debug)

return M
