----------------
--- MONITORS ---
----------------

hl.monitor({
    output = "eDP-1",
    mode = "preferred",
    position = "0x0",
    scale = "1",
})

hl.monitor({
    output = "HDMI-A-1",
    mode = "preferred",
    position = "0x-1080",
    scale = "1",
})

-------------------
--- MY PROGRAMS ---
-------------------

local terminal = "kitty"
local menu = "rofi -show-icons -show drun"
local file_manager = "nautilus --new-window"
local browser = "firefox"

-----------------
--- AUTOSTART ---
-----------------

hl.on("hyprland.start", function()
    hl.exec_cmd("exec /usr/lib/polkit-gnome/polkit-gnome-authentication-agent-1")

    hl.exec_cmd("waybar")
    hl.exec_cmd("hyprpaper")
    hl.exec_cmd("hypridle")

    hl.exec_cmd("wl-paste --type text --watch cliphist store")
    hl.exec_cmd("wl-paste --type image --watch cliphist store")
    hl.exec_cmd("wl-clip-persist --clipboard regular")

    hl.exec_cmd("rfkill unblock bluetooth")
    hl.exec_cmd("bluetoothctl power off")
end)

-----------------------------
--- ENVIRONMENT VARIABLES ---
-----------------------------

hl.env("XCURSOR_SIZE", "24")
hl.env("HYPRCURSOR_SIZE", "24")

---------------------
--- LOOK AND FEEL ---
---------------------

hl.config({
    general = {
        gaps_in = 0,
        gaps_out = 0,

        border_size = 0,

        allow_tearing = false,

        layout = "dwindle",
    },

    decoration = {
        inactive_opacity = 0.85,
        shadow = { enabled = false },
        blur = {
            enabled = true,
            passes = 4,
        },
    },

    animations = { enabled = true },
})

hl.animation({
    leaf = "global",
    enabled = true,
    speed = 3,
    bezier = "default",
})

hl.config({
    dwindle = {
        preserve_split = true,
    },
})

----------------
----  MISC  ----
----------------

hl.config({
    misc = {
        force_default_wallpaper = -1,
        disable_hyprland_logo = true,
    },
})

-------------
--- INPUT ---
-------------

hl.config({
    input = {
        kb_layout = "us, ara",
        kb_options = "grp:alts_toggle",

        follow_mouse = 2,
        sensitivity = 0,

        touchpad = {
            natural_scroll = true,
        },

        tablet = {
            output = "eDP-1",
        },
    },
})

--------------------------------
--- KEYBINDINGS AND GESTURES ---
--------------------------------

local mainMod = "SUPER"

-- Workspaces and Windows
hl.bind(mainMod .. " + backspace", hl.dsp.window.close())
hl.bind(mainMod .. " + f", hl.dsp.window.fullscreen())
hl.bind(mainMod .. " + u", hl.dsp.layout("togglesplit"))
hl.bind(mainMod .. " + i", hl.dsp.window.pseudo())
hl.bind(mainMod .. " + backslash", hl.dsp.window.float({ action = "toggle" }))
hl.bind(mainMod .. " + p", hl.dsp.window.pin())

for i = 1, 5 do
    hl.bind(mainMod .. " + " .. i, hl.dsp.focus({ workspace = i }))
    hl.bind(mainMod .. " + SHIFT + " .. i, hl.dsp.window.move({ workspace = i }))
end

hl.bind(mainMod .. " + 0", hl.dsp.workspace.toggle_special("background"))
hl.bind(mainMod .. " + SHIFT + 0", hl.dsp.window.move({ workspace = "special:background" }))

hl.gesture({ fingers = 3, direction = "horizontal", action = "workspace" })

hl.bind(mainMod .. " + h", hl.dsp.focus({ direction = "left" }))
hl.bind(mainMod .. " + j", hl.dsp.focus({ direction = "down" }))
hl.bind(mainMod .. " + k", hl.dsp.focus({ direction = "up" }))
hl.bind(mainMod .. " + l", hl.dsp.focus({ direction = "right" }))

hl.bind(mainMod .. " + left", hl.dsp.focus({ direction = "left" }))
hl.bind(mainMod .. " + down", hl.dsp.focus({ direction = "down" }))
hl.bind(mainMod .. " + up", hl.dsp.focus({ direction = "up" }))
hl.bind(mainMod .. " + right", hl.dsp.focus({ direction = "right" }))

hl.bind(mainMod .. " + tab", hl.dsp.window.cycle_next())

hl.bind(mainMod .. " + semicolon", hl.dsp.window.resize({ x = "-40", y = "0", relative = true }),
    { locked = true, repeating = true })
hl.bind(mainMod .. " + SHIFT + semicolon", hl.dsp.window.resize({ x = "0", y = "40", relative = true }),
    { locked = true, repeating = true })
hl.bind(mainMod .. " + apostrophe", hl.dsp.window.resize({ x = "40", y = "0", relative = true }),
    { locked = true, repeating = true })
hl.bind(mainMod .. " + SHIFT + apostrophe", hl.dsp.window.resize({ x = "0", y = "-40", relative = true }),
    { locked = true, repeating = true })


hl.bind(mainMod .. " + mouse:273", hl.dsp.window.resize(),
    { mouse = true })

hl.bind(mainMod .. " + bracketleft", hl.dsp.window.swap({ direction = "left" }))
hl.bind(mainMod .. " + SHIFT + bracketleft", hl.dsp.window.swap({ direction = "down" }))
hl.bind(mainMod .. " + bracketright", hl.dsp.window.swap({ direction = "right" }))
hl.bind(mainMod .. " + SHIFT + bracketright", hl.dsp.window.swap({ direction = "up" }))

hl.bind(mainMod .. " + mouse:272", hl.dsp.window.drag(),
    { mouse = true })

-- XF86 Buttons
hl.bind("XF86MonBrightnessUp", hl.dsp.exec_cmd("brightnessctl s 5%+"),
    { locked = true, repeating = true })
hl.bind("XF86MonBrightnessDown", hl.dsp.exec_cmd("brightnessctl s 5%-"),
    { locked = true, repeating = true })
hl.bind("XF86AudioRaiseVolume", hl.dsp.exec_cmd("wpctl set-volume -l 1 @DEFAULT_AUDIO_SINK@ 5%+"),
    { locked = true, repeating = true })
hl.bind("XF86AudioLowerVolume", hl.dsp.exec_cmd("wpctl set-volume @DEFAULT_AUDIO_SINK@ 5%-"),
    { locked = true, repeating = true })
hl.bind("XF86AudioMute", hl.dsp.exec_cmd("wpctl set-mute @DEFAULT_AUDIO_SINK@ toggle"),
    { locked = true, repeating = true })
hl.bind("XF86AudioMicMute", hl.dsp.exec_cmd("wpctl set-mute @DEFAULT_AUDIO_SOURCE@ toggle"),
    { locked = true, repeating = true })
hl.bind("XF86AudioNext", hl.dsp.exec_cmd("playerctl next"),
    { locked = true, repeating = true })
hl.bind("XF86AudioPause", hl.dsp.exec_cmd("playerctl play-pause"),
    { locked = true, repeating = true })
hl.bind("XF86AudioPlay", hl.dsp.exec_cmd("playerctl play-pause"),
    { locked = true, repeating = true })
hl.bind("XF86AudioPrev", hl.dsp.exec_cmd("playerctl previous"),
    { locked = true, repeating = true })

-- Screenshot
hl.bind("print", hl.dsp.exec_cmd("hyprshot screen"))
hl.bind(mainMod .. " + s", hl.dsp.exec_cmd("hyprshot region"))

-- Apps and Commands
hl.bind(mainMod .. " + delete", hl.dsp.exec_cmd("hyprlock"))

hl.bind(mainMod .. " + o", hl.dsp.exec_cmd(menu))
hl.bind(mainMod .. " + v", hl.dsp.exec_cmd("cliphist list | rofi -dmenu | cliphist decode | wl-copy"))
hl.bind(mainMod .. " + e", hl.dsp.exec_cmd("rofi -modi emoji -show emoji"))

hl.bind(mainMod .. " + space", hl.dsp.exec_cmd("playerctl play-pause"),
    { locked = true, repeating = true })

hl.bind(mainMod .. " + return", hl.dsp.exec_cmd(terminal))
hl.bind(mainMod .. " + period", hl.dsp.exec_cmd(file_manager))
hl.bind(mainMod .. " + slash", hl.dsp.exec_cmd(browser))

-- Submaps
hl.bind(mainMod .. " + grave", hl.dsp.submap("Control"))
hl.define_submap("Control", function()
    hl.bind("escape", hl.dsp.submap("reset"))
    hl.bind("h", hl.dsp.exec_cmd("wpctl set-volume -l 1 @DEFAULT_AUDIO_SINK@ 5%-"),
        { locked = true, repeating = true })
    hl.bind("SHIFT + h", hl.dsp.exec_cmd("wpctl set-volume -l 1 @DEFAULT_AUDIO_SINK@ 10%-"),
        { locked = true, repeating = true })
    hl.bind("j", hl.dsp.exec_cmd("brightnessctl s 5%-"),
        { locked = true, repeating = true })
    hl.bind("SHIFT + j", hl.dsp.exec_cmd("brightnessctl s 10%-"),
        { locked = true, repeating = true })
    hl.bind("k", hl.dsp.exec_cmd("brightnessctl s 5%+"),
        { locked = true, repeating = true })
    hl.bind("SHIFT + k", hl.dsp.exec_cmd("brightnessctl s 10%+"),
        { locked = true, repeating = true })
    hl.bind("l", hl.dsp.exec_cmd("wpctl set-volume -l 1 @DEFAULT_AUDIO_SINK@ 5%+"),
        { locked = true, repeating = true })
    hl.bind("SHIFT + l", hl.dsp.exec_cmd("wpctl set-volume -l 1 @DEFAULT_AUDIO_SINK@ 10%+"),
        { locked = true, repeating = true })
end)


------------------------------
--- WINDOWS AND WORKSPACES ---
------------------------------

hl.workspace_rule({ workspace = "1", default = true, monitor = "eDP-1" })
hl.workspace_rule({ workspace = "2", monitor = "eDP-1" })
hl.workspace_rule({ workspace = "3", monitor = "eDP-1" })
hl.workspace_rule({ workspace = "4", monitor = "eDP-1" })
hl.workspace_rule({ workspace = "5", monitor = "HDMI-A-1" })

local float_clients = {
    "Calculator",
}

local pin_clients = {
    "Picture-in-Picture",
    "Picture in picture",
}

for index, value in ipairs(float_clients) do
    hl.window_rule({ match = { class = value }, float = true })
end

for index, value in ipairs(pin_clients) do
    hl.window_rule({ match = { title = value }, float = true })
    hl.window_rule({ match = { title = value }, pin = true })
end

hl.window_rule({
    name = "suppress-maximize-events",
    match = { class = ".*" },
    suppress_event = "maximize",
})

hl.window_rule({
    name = "fix-xwayland-drags",
    match = {
        class = "^$",
        title = "^$",
        xwayland = true,
        float = true,
        fullscreen = false,
        pin = false,
    },
    no_focus = true,
})

hl.window_rule({
    name = "move-hyprland-run",
    match = { class = "hyprland-run" },
    move = "20 monitor_h-120",
    float = true,
})
