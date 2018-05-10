if not nyagos then
    print("This is a script for nyagos not lua.exe")
    os.exit()
end

nyagos.alias.cd = function(args)
  local target = nil
  if #args == 0 then
    target = {}
  else
    local parent, pattern = args[1]:match('(.*[\\/])(.*)')
    parent = parent or ''
    pattern = pattern or args[1]
    -- except '..' and drive
    if pattern == '..' or pattern:find('^[%u%l]:$') then
      target = {parent .. pattern}
    elseif pattern == '' then
      target = {parent}
    else
      local dirs = getDirs(parent, pattern)
      if #dirs == 0 then
        print(pattern .. ': No such directory')
      elseif #dirs == 1 then
        target = {dirs[1]}
      else
        local d = nyagos.box(dirs)
        print()
        if #d > 0 then
          target = {d}
        end
      end
    end
  end
  if target ~= nil then
    _cd(target)
  end
end

getDirs = function(parent, pattern)
  if parent:find('^[\\/].*') then
    if not(parent:find('^[\\/][\\/]+')) then
      parent = nyagos.getwd():gsub('[\\/].*', '') .. parent
    end
  end
  local line = nyagos.eval('ls -la ' .. parent) -- `ls` which is embedded in nyagos
  local complst = {}
  for i, e in ipairs(split(line, '[\r\n]')) do
    local t = tostring(e:gsub('.-%s+', '', 5))
    if t ~= './' and t ~= '../' then
      table.insert(complst, '"' .. t .. '"')
    end
  end

  local opt = '-f ' -- forword match
  -- if the first char is lower, ignorecase
  if pattern:sub(0):match('%l') then
    opt = opt .. '-i '
  end

  local dirs = {}
  -- is max length of command line 8192 ?
  local n = 1000
  for i = 1, #complst, n do
    local lst = {}
    for j = 1, n do
      lst[j] = complst[j + (i-1)]
    end
    for _,e in pairs(split(nyagos.eval('gmgmgm -f ' .. opt .. pattern .. ' ' .. table.concat(lst, ' ')), '[\r\n]')) do
      if e:match('.*/$') then
        table.insert(dirs, #dirs+1, parent .. e:gsub('/$', ''))
      end
    end
  end
  return dirs
end

_cd = function(args)
  local t = ''
  if #args >= 1 then
    t = '"' .. args[1]:gsub('\\', '/') .. '"'
  end
  r, err = nyagos.exec('__cd__ ' .. t)
  return r, err
end

split = function(str, pat)
  local t = {}  -- NOTE: use {n = 0} in Lua-5.0
  local fpat = '(.-)' .. pat
  local last_end = 1
  local s, e, cap = str:find(fpat, 1)
  while s do
    if s ~= 1 or cap ~= '' then
      table.insert(t, cap)
    end
    last_end = e+1
    s, e, cap = str:find(fpat, last_end)
  end
  if last_end <= #str then
    cap = str:sub(last_end)
    table.insert(t, cap)
  end
  return t
end
