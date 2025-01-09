
local M = {}

-- MATH

function M.random_seed()

	math.randomseed(os.time())
	for i = 1, 5 do
		math.random()
	end

end


function M.get_angle(quat)

	local angle = select(3, vmath.quat_to_euler(quat))
	return math.rad(angle)

end

-- VECTORS

function M.random_in_circle(radius, center_offset, border_offset)
	
	local r = (radius - border_offset - center_offset) * math.sqrt(math.random()) + center_offset
	local a = math.random() * 2 * math.pi
	local x = r * math.cos(a)
	local y = r * math.sin(a)
	return vmath.vector3(x, y, 0)
	
end

-- GUI

function M.set_enabled(node)

	gui.set_enabled(gui.get_node(node), true)

end


function M.set_disabled(node)

	gui.set_enabled(gui.get_node(node), false)

end


function M.set_text(node, text, color, time)

	local color_old = gui.get_color(gui.get_node(node))
	local color_new = color and vmath.vector3(color[1], color[2], color[3]) or color_old
	gui.set_text(gui.get_node(node), text)
	gui.set_color(gui.get_node(node), color_new)
	if time then
		timer.delay(time, false, function()
			gui.set_color(gui.get_node(node), color_old)
		end)
	end
	
end

return M
