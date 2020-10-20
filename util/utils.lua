local M = {}



function M.wobble(node, scale)
	gui.cancel_animation(node, "scale.x")
	gui.cancel_animation(node, "scale.y")
	gui.set_scale(node, scale or vmath.vector3(1))
	gui.animate(node, "scale.x", 1.1, gui.EASING_INOUTQUAD, 2.5, 0, nil, gui.PLAYBACK_LOOP_PINGPONG)
	gui.animate(node, "scale.y", 1.15, gui.EASING_INOUTQUAD, 2.5, 0.75, nil, gui.PLAYBACK_LOOP_PINGPONG)
end

function M.stop_wobble(node, scale)
	gui.cancel_animation(node, "scale.x")
	gui.cancel_animation(node, "scale.y")
	gui.set_scale(node, scale or vmath.vector3(1))
end


return M