
local M = {}

function M.sign(number)
	return number > 0 and 1 or (number == 0 and 0 or -1)
end


function M.rotate(from, to)
	if tonumber(to) then
		if to == nil then
			to = go.get_id()
		end
		go.set(from, "euler.z", to)
	else
		local direction = go.get_position(to) - go.get_position(from)
		local rotation = vmath.quat_rotation_z(math.atan2(direction.y, direction.x))
		go.set_rotation(rotation, from)
	end
end


function M.random_seed()
	math.randomseed(os.time())
	for i = 1, 5 do
		math.random()
	end
end


function M.random_in_circle(radius, center_offset, border_offset)
	local r = (radius - border_offset - center_offset) * math.sqrt(math.random()) + center_offset
	local a = math.random() * 2 * math.pi
	local x = r * math.cos(a)
	local y = r * math.sin(a)
	return vmath.vector3(x, y, 0)
end

return M
