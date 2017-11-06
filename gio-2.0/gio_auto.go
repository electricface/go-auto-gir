package gio

/*
#cgo pkg-config: gio-2.0 gio-unix-2.0
#include <gio/gdesktopappinfo.h>
#include <gio/gfiledescriptorbased.h>
#include <gio/gio.h>
#include <gio/gunixconnection.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixfdmessage.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixmounts.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixsocketaddress.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"
import "github.com/electricface/go-auto-gir/gobject-2.0"
import "github.com/electricface/go-auto-gir/glib-2.0"

// Interface AppInfo
type AppInfo struct {
	Ptr unsafe.Pointer
}

func (v AppInfo) native() *C.GAppInfo {
	return (*C.GAppInfo)(v.Ptr)
}
func wrapAppInfo(p *C.GAppInfo) AppInfo {
	return AppInfo{unsafe.Pointer(p)}
}
func WrapAppInfo(p unsafe.Pointer) AppInfo {
	return AppInfo{p}
}

// GetId is a wrapper around g_app_info_get_id().
func (appinfo AppInfo) GetId() string {
	ret0 := C.g_app_info_get_id(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetSupportedTypes is a wrapper around g_app_info_get_supported_types().
func (appinfo AppInfo) GetSupportedTypes() []string {
	ret0 := C.g_app_info_get_supported_types(appinfo.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
	}
	return ret
}

// Object DesktopAppInfo
type DesktopAppInfo struct {
	gobject.Object
}

func (v DesktopAppInfo) native() *C.GDesktopAppInfo {
	return (*C.GDesktopAppInfo)(v.Ptr)
}
func wrapDesktopAppInfo(p *C.GDesktopAppInfo) (v DesktopAppInfo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDesktopAppInfo(p unsafe.Pointer) (v DesktopAppInfo) {
	v.Ptr = p
	return
}
func (v DesktopAppInfo) AppInfo() AppInfo {
	return WrapAppInfo(v.Ptr)
}

// DesktopAppInfoNew is a wrapper around g_desktop_app_info_new().
func DesktopAppInfoNew(desktop_id string) DesktopAppInfo {
	desktop_id0 := C.CString(desktop_id)
	ret0 := C.g_desktop_app_info_new(desktop_id0)
	C.free(unsafe.Pointer(desktop_id0)) /*ch:<stdlib.h>*/
	return wrapDesktopAppInfo(ret0)
}

// DesktopAppInfoNewFromFilename is a wrapper around g_desktop_app_info_new_from_filename().
func DesktopAppInfoNewFromFilename(filename string) DesktopAppInfo {
	filename0 := C.CString(filename)
	ret0 := C.g_desktop_app_info_new_from_filename(filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	return wrapDesktopAppInfo(ret0)
}

// GetActionName is a wrapper around g_desktop_app_info_get_action_name().
func (info DesktopAppInfo) GetActionName(action_name string) string {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_desktop_app_info_get_action_name(info.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetBoolean is a wrapper around g_desktop_app_info_get_boolean().
func (info DesktopAppInfo) GetBoolean(key string) bool {
	key0 := C.CString(key)
	ret0 := C.g_desktop_app_info_get_boolean(info.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCategories is a wrapper around g_desktop_app_info_get_categories().
func (info DesktopAppInfo) GetCategories() string {
	ret0 := C.g_desktop_app_info_get_categories(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetFilename is a wrapper around g_desktop_app_info_get_filename().
func (info DesktopAppInfo) GetFilename() string {
	ret0 := C.g_desktop_app_info_get_filename(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetGenericName is a wrapper around g_desktop_app_info_get_generic_name().
func (info DesktopAppInfo) GetGenericName() string {
	ret0 := C.g_desktop_app_info_get_generic_name(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetIsHidden is a wrapper around g_desktop_app_info_get_is_hidden().
func (info DesktopAppInfo) GetIsHidden() bool {
	ret0 := C.g_desktop_app_info_get_is_hidden(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetNodisplay is a wrapper around g_desktop_app_info_get_nodisplay().
func (info DesktopAppInfo) GetNodisplay() bool {
	ret0 := C.g_desktop_app_info_get_nodisplay(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowIn is a wrapper around g_desktop_app_info_get_show_in().
func (info DesktopAppInfo) GetShowIn(desktop_env string) bool {
	desktop_env0 := (*C.gchar)(C.CString(desktop_env))
	ret0 := C.g_desktop_app_info_get_show_in(info.native(), desktop_env0)
	C.free(unsafe.Pointer(desktop_env0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// GetStartupWmClass is a wrapper around g_desktop_app_info_get_startup_wm_class().
func (info DesktopAppInfo) GetStartupWmClass() string {
	ret0 := C.g_desktop_app_info_get_startup_wm_class(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetString is a wrapper around g_desktop_app_info_get_string().
func (info DesktopAppInfo) GetString(key string) string {
	key0 := C.CString(key)
	ret0 := C.g_desktop_app_info_get_string(info.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasKey is a wrapper around g_desktop_app_info_has_key().
func (info DesktopAppInfo) HasKey(key string) bool {
	key0 := C.CString(key)
	ret0 := C.g_desktop_app_info_has_key(info.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LaunchAction is a wrapper around g_desktop_app_info_launch_action().
func (info DesktopAppInfo) LaunchAction(action_name string, launch_context AppLaunchContext) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_desktop_app_info_launch_action(info.native(), action_name0, launch_context.native())
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ListActions is a wrapper around g_desktop_app_info_list_actions().
func (info DesktopAppInfo) ListActions() []string {
	ret0 := C.g_desktop_app_info_list_actions(info.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	return ret
}

// Object AppLaunchContext
type AppLaunchContext struct {
	gobject.Object
}

func (v AppLaunchContext) native() *C.GAppLaunchContext {
	return (*C.GAppLaunchContext)(v.Ptr)
}
func wrapAppLaunchContext(p *C.GAppLaunchContext) (v AppLaunchContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppLaunchContext(p unsafe.Pointer) (v AppLaunchContext) {
	v.Ptr = p
	return
}

// AppLaunchContextNew is a wrapper around g_app_launch_context_new().
func AppLaunchContextNew() AppLaunchContext {
	ret0 := C.g_app_launch_context_new()
	return wrapAppLaunchContext(ret0)
}

// GetEnvironment is a wrapper around g_app_launch_context_get_environment().
func (context AppLaunchContext) GetEnvironment() []string {
	ret0 := C.g_app_launch_context_get_environment(context.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// LaunchFailed is a wrapper around g_app_launch_context_launch_failed().
func (context AppLaunchContext) LaunchFailed(startup_notify_id string) {
	startup_notify_id0 := C.CString(startup_notify_id)
	C.g_app_launch_context_launch_failed(context.native(), startup_notify_id0)
	C.free(unsafe.Pointer(startup_notify_id0)) /*ch:<stdlib.h>*/
}

// Setenv is a wrapper around g_app_launch_context_setenv().
func (context AppLaunchContext) Setenv(variable string, value string) {
	variable0 := C.CString(variable)
	value0 := C.CString(value)
	C.g_app_launch_context_setenv(context.native(), variable0, value0)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))    /*ch:<stdlib.h>*/
}

// Unsetenv is a wrapper around g_app_launch_context_unsetenv().
func (context AppLaunchContext) Unsetenv(variable string) {
	variable0 := C.CString(variable)
	C.g_app_launch_context_unsetenv(context.native(), variable0)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
}

// Object Settings
type Settings struct {
	gobject.Object
}

func (v Settings) native() *C.GSettings {
	return (*C.GSettings)(v.Ptr)
}
func wrapSettings(p *C.GSettings) (v Settings) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSettings(p unsafe.Pointer) (v Settings) {
	v.Ptr = p
	return
}

// SettingsNew is a wrapper around g_settings_new().
func SettingsNew(schema_id string) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	ret0 := C.g_settings_new(schema_id0)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewFull is a wrapper around g_settings_new_full().
func SettingsNewFull(schema SettingsSchema, backend SettingsBackend, path string) Settings {
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_settings_new_full(schema.native(), backend.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewWithBackend is a wrapper around g_settings_new_with_backend().
func SettingsNewWithBackend(schema_id string, backend SettingsBackend) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	ret0 := C.g_settings_new_with_backend(schema_id0, backend.native())
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewWithBackendAndPath is a wrapper around g_settings_new_with_backend_and_path().
func SettingsNewWithBackendAndPath(schema_id string, backend SettingsBackend, path string) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_settings_new_with_backend_and_path(schema_id0, backend.native(), path0)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(path0))      /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewWithPath is a wrapper around g_settings_new_with_path().
func SettingsNewWithPath(schema_id string, path string) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_settings_new_with_path(schema_id0, path0)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(path0))      /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// Apply is a wrapper around g_settings_apply().
func (settings Settings) Apply() {
	C.g_settings_apply(settings.native())
}

// Bind is a wrapper around g_settings_bind().
func (settings Settings) Bind(key string, object gobject.Object, property string, flags SettingsBindFlags) {
	key0 := (*C.gchar)(C.CString(key))
	property0 := (*C.gchar)(C.CString(property))
	C.g_settings_bind(settings.native(), key0, C.gpointer((C.gpointer)(object.Ptr)), property0, C.GSettingsBindFlags(flags))
	C.free(unsafe.Pointer(key0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// BindWritable is a wrapper around g_settings_bind_writable().
func (settings Settings) BindWritable(key string, object gobject.Object, property string, inverted bool) {
	key0 := (*C.gchar)(C.CString(key))
	property0 := (*C.gchar)(C.CString(property))
	C.g_settings_bind_writable(settings.native(), key0, C.gpointer((C.gpointer)(object.Ptr)), property0, C.gboolean(util.Bool2Int(inverted)) /*go:.util*/)
	C.free(unsafe.Pointer(key0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// CreateAction is a wrapper around g_settings_create_action().
func (settings Settings) CreateAction(key string) Action {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_create_action(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return wrapAction(ret0)
}

// Delay is a wrapper around g_settings_delay().
func (settings Settings) Delay() {
	C.g_settings_delay(settings.native())
}

// GetBoolean is a wrapper around g_settings_get_boolean().
func (settings Settings) GetBoolean(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_boolean(settings.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChild is a wrapper around g_settings_get_child().
func (settings Settings) GetChild(name string) Settings {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_settings_get_child(settings.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// GetDefaultValue is a wrapper around g_settings_get_default_value().
func (settings Settings) GetDefaultValue(key string) glib.Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_default_value(settings.native(), key0)
	C.free(unsafe.Pointer(key0))                  /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetDouble is a wrapper around g_settings_get_double().
func (settings Settings) GetDouble(key string) float64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_double(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetEnum is a wrapper around g_settings_get_enum().
func (settings Settings) GetEnum(key string) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_enum(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetFlags is a wrapper around g_settings_get_flags().
func (settings Settings) GetFlags(key string) uint {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_flags(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// GetHasUnapplied is a wrapper around g_settings_get_has_unapplied().
func (settings Settings) GetHasUnapplied() bool {
	ret0 := C.g_settings_get_has_unapplied(settings.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInt is a wrapper around g_settings_get_int().
func (settings Settings) GetInt(key string) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_int(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetInt64 is a wrapper around g_settings_get_int64().
func (settings Settings) GetInt64(key string) int64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_int64(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int64(ret0)
}

// GetString is a wrapper around g_settings_get_string().
func (settings Settings) GetString(key string) string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_string(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetStrv is a wrapper around g_settings_get_strv().
func (settings Settings) GetStrv(key string) []string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_strv(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUint is a wrapper around g_settings_get_uint().
func (settings Settings) GetUint(key string) uint {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_uint(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// GetUint64 is a wrapper around g_settings_get_uint64().
func (settings Settings) GetUint64(key string) uint64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_uint64(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint64(ret0)
}

// GetUserValue is a wrapper around g_settings_get_user_value().
func (settings Settings) GetUserValue(key string) glib.Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_user_value(settings.native(), key0)
	C.free(unsafe.Pointer(key0))                  /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetValue is a wrapper around g_settings_get_value().
func (settings Settings) GetValue(key string) glib.Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_value(settings.native(), key0)
	C.free(unsafe.Pointer(key0))                  /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// IsWritable is a wrapper around g_settings_is_writable().
func (settings Settings) IsWritable(name string) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_settings_is_writable(settings.native(), name0)
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListChildren is a wrapper around g_settings_list_children().
func (settings Settings) ListChildren() []string {
	ret0 := C.g_settings_list_children(settings.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// Reset is a wrapper around g_settings_reset().
func (settings Settings) Reset(key string) {
	key0 := (*C.gchar)(C.CString(key))
	C.g_settings_reset(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// Revert is a wrapper around g_settings_revert().
func (settings Settings) Revert() {
	C.g_settings_revert(settings.native())
}

// SetBoolean is a wrapper around g_settings_set_boolean().
func (settings Settings) SetBoolean(key string, value bool) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_boolean(settings.native(), key0, C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetDouble is a wrapper around g_settings_set_double().
func (settings Settings) SetDouble(key string, value float64) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_double(settings.native(), key0, C.gdouble(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetEnum is a wrapper around g_settings_set_enum().
func (settings Settings) SetEnum(key string, value int) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_enum(settings.native(), key0, C.gint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetFlags is a wrapper around g_settings_set_flags().
func (settings Settings) SetFlags(key string, value uint) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_flags(settings.native(), key0, C.guint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetInt is a wrapper around g_settings_set_int().
func (settings Settings) SetInt(key string, value int) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_int(settings.native(), key0, C.gint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetInt64 is a wrapper around g_settings_set_int64().
func (settings Settings) SetInt64(key string, value int64) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_int64(settings.native(), key0, C.gint64(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetString is a wrapper around g_settings_set_string().
func (settings Settings) SetString(key string, value string) bool {
	key0 := (*C.gchar)(C.CString(key))
	value0 := (*C.gchar)(C.CString(value))
	ret0 := C.g_settings_set_string(settings.native(), key0, value0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))  /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetStrv is a wrapper around g_settings_set_strv().
func (settings Settings) SetStrv(key string, value []string) bool {
	key0 := (*C.gchar)(C.CString(key))
	value0 := make([]*C.gchar, len(value))
	for idx, elemG := range value {
		elem := (*C.gchar)(C.CString(elemG))
		value0[idx] = elem
	}
	var value0Ptr **C.gchar
	if len(value0) > 0 {
		value0Ptr = &value0[0]
	}
	ret0 := C.g_settings_set_strv(settings.native(), key0, value0Ptr)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	for _, elem := range value0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUint is a wrapper around g_settings_set_uint().
func (settings Settings) SetUint(key string, value uint) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_uint(settings.native(), key0, C.guint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUint64 is a wrapper around g_settings_set_uint64().
func (settings Settings) SetUint64(key string, value uint64) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_uint64(settings.native(), key0, C.guint64(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetValue is a wrapper around g_settings_set_value().
func (settings Settings) SetValue(key string, value glib.Variant) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_value(settings.native(), key0, (*C.GVariant)(value.Ptr))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SettingsSync is a wrapper around g_settings_sync().
func SettingsSync() {
	C.g_settings_sync()
}

// SettingsUnbind is a wrapper around g_settings_unbind().
func SettingsUnbind(object gobject.Object, property string) {
	property0 := (*C.gchar)(C.CString(property))
	C.g_settings_unbind(C.gpointer((C.gpointer)(object.Ptr)), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// Struct SettingsSchema
type SettingsSchema struct {
	Ptr unsafe.Pointer
}

func (v SettingsSchema) native() *C.GSettingsSchema {
	return (*C.GSettingsSchema)(v.Ptr)
}
func wrapSettingsSchema(p *C.GSettingsSchema) SettingsSchema {
	return SettingsSchema{unsafe.Pointer(p)}
}
func WrapSettingsSchema(p unsafe.Pointer) SettingsSchema {
	return SettingsSchema{p}
}

// Struct SettingsBackend
type SettingsBackend struct {
	Ptr unsafe.Pointer
}

func (v SettingsBackend) native() *C.GSettingsBackend {
	return (*C.GSettingsBackend)(v.Ptr)
}
func wrapSettingsBackend(p *C.GSettingsBackend) SettingsBackend {
	return SettingsBackend{unsafe.Pointer(p)}
}
func WrapSettingsBackend(p unsafe.Pointer) SettingsBackend {
	return SettingsBackend{p}
}

// Interface Action
type Action struct {
	Ptr unsafe.Pointer
}

func (v Action) native() *C.GAction {
	return (*C.GAction)(v.Ptr)
}
func wrapAction(p *C.GAction) Action {
	return Action{unsafe.Pointer(p)}
}
func WrapAction(p unsafe.Pointer) Action {
	return Action{p}
}

// Interface File
type File struct {
	Ptr unsafe.Pointer
}

func (v File) native() *C.GFile {
	return (*C.GFile)(v.Ptr)
}
func wrapFile(p *C.GFile) File {
	return File{unsafe.Pointer(p)}
}
func WrapFile(p unsafe.Pointer) File {
	return File{p}
}

// AppendTo is a wrapper around g_file_append_to().
func (file File) AppendTo(flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_append_to(file.native(), C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// AppendToFinish is a wrapper around g_file_append_to_finish().
func (file File) AppendToFinish(res AsyncResult) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_append_to_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// CopyAttributes is a wrapper around g_file_copy_attributes().
func (source File) CopyAttributes(destination File, flags FileCopyFlags, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_copy_attributes(source.native(), destination.native(), C.GFileCopyFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CopyFinish is a wrapper around g_file_copy_finish().
func (file File) CopyFinish(res AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_copy_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Create is a wrapper around g_file_create().
func (file File) Create(flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_create(file.native(), C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// CreateFinish is a wrapper around g_file_create_finish().
func (file File) CreateFinish(res AsyncResult) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_create_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// CreateReadwrite is a wrapper around g_file_create_readwrite().
func (file File) CreateReadwrite(flags FileCreateFlags, cancellable Cancellable) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_create_readwrite(file.native(), C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// CreateReadwriteFinish is a wrapper around g_file_create_readwrite_finish().
func (file File) CreateReadwriteFinish(res AsyncResult) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_create_readwrite_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// Delete is a wrapper around g_file_delete().
func (file File) Delete(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_delete(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// DeleteFinish is a wrapper around g_file_delete_finish().
func (file File) DeleteFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_delete_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Dup is a wrapper around g_file_dup().
func (file File) Dup() File {
	ret0 := C.g_file_dup(file.native())
	return wrapFile(ret0)
}

// EjectMountableWithOperationFinish is a wrapper around g_file_eject_mountable_with_operation_finish().
func (file File) EjectMountableWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_eject_mountable_with_operation_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EnumerateChildrenFinish is a wrapper around g_file_enumerate_children_finish().
func (file File) EnumerateChildrenFinish(res AsyncResult) (FileEnumerator, error) {
	var err glib.Error
	ret0 := C.g_file_enumerate_children_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileEnumerator{}, err.GoValue()
	}
	return wrapFileEnumerator(ret0), nil
}

// Equal is a wrapper around g_file_equal().
func (file1 File) Equal(file2 File) bool {
	ret0 := C.g_file_equal(file1.native(), file2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// FindEnclosingMountFinish is a wrapper around g_file_find_enclosing_mount_finish().
func (file File) FindEnclosingMountFinish(res AsyncResult) (Mount, error) {
	var err glib.Error
	ret0 := C.g_file_find_enclosing_mount_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Mount{}, err.GoValue()
	}
	return wrapMount(ret0), nil
}

// GetBasename is a wrapper around g_file_get_basename().
func (file File) GetBasename() string {
	ret0 := C.g_file_get_basename(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetChild is a wrapper around g_file_get_child().
func (file File) GetChild(name string) File {
	name0 := C.CString(name)
	ret0 := C.g_file_get_child(file.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// GetChildForDisplayName is a wrapper around g_file_get_child_for_display_name().
func (file File) GetChildForDisplayName(display_name string) (File, error) {
	display_name0 := C.CString(display_name)
	var err glib.Error
	ret0 := C.g_file_get_child_for_display_name(file.native(), display_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// GetParent is a wrapper around g_file_get_parent().
func (file File) GetParent() File {
	ret0 := C.g_file_get_parent(file.native())
	return wrapFile(ret0)
}

// GetParseName is a wrapper around g_file_get_parse_name().
func (file File) GetParseName() string {
	ret0 := C.g_file_get_parse_name(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetPath is a wrapper around g_file_get_path().
func (file File) GetPath() string {
	ret0 := C.g_file_get_path(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetRelativePath is a wrapper around g_file_get_relative_path().
func (parent File) GetRelativePath(descendant File) string {
	ret0 := C.g_file_get_relative_path(parent.native(), descendant.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUri is a wrapper around g_file_get_uri().
func (file File) GetUri() string {
	ret0 := C.g_file_get_uri(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUriScheme is a wrapper around g_file_get_uri_scheme().
func (file File) GetUriScheme() string {
	ret0 := C.g_file_get_uri_scheme(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasParent is a wrapper around g_file_has_parent().
func (file File) HasParent(parent File) bool {
	ret0 := C.g_file_has_parent(file.native(), parent.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasPrefix is a wrapper around g_file_has_prefix().
func (file File) HasPrefix(prefix File) bool {
	ret0 := C.g_file_has_prefix(file.native(), prefix.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasUriScheme is a wrapper around g_file_has_uri_scheme().
func (file File) HasUriScheme(uri_scheme string) bool {
	uri_scheme0 := C.CString(uri_scheme)
	ret0 := C.g_file_has_uri_scheme(file.native(), uri_scheme0)
	C.free(unsafe.Pointer(uri_scheme0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// IsNative is a wrapper around g_file_is_native().
func (file File) IsNative() bool {
	ret0 := C.g_file_is_native(file.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MakeDirectory is a wrapper around g_file_make_directory().
func (file File) MakeDirectory(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_make_directory(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MakeDirectoryFinish is a wrapper around g_file_make_directory_finish().
func (file File) MakeDirectoryFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_make_directory_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MakeDirectoryWithParents is a wrapper around g_file_make_directory_with_parents().
func (file File) MakeDirectoryWithParents(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_make_directory_with_parents(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MakeSymbolicLink is a wrapper around g_file_make_symbolic_link().
func (file File) MakeSymbolicLink(symlink_value string, cancellable Cancellable) (bool, error) {
	symlink_value0 := C.CString(symlink_value)
	var err glib.Error
	ret0 := C.g_file_make_symbolic_link(file.native(), symlink_value0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(symlink_value0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Monitor is a wrapper around g_file_monitor().
func (file File) Monitor(flags FileMonitorFlags, cancellable Cancellable) (FileMonitor, error) {
	var err glib.Error
	ret0 := C.g_file_monitor(file.native(), C.GFileMonitorFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileMonitor{}, err.GoValue()
	}
	return wrapFileMonitor(ret0), nil
}

// MonitorDirectory is a wrapper around g_file_monitor_directory().
func (file File) MonitorDirectory(flags FileMonitorFlags, cancellable Cancellable) (FileMonitor, error) {
	var err glib.Error
	ret0 := C.g_file_monitor_directory(file.native(), C.GFileMonitorFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileMonitor{}, err.GoValue()
	}
	return wrapFileMonitor(ret0), nil
}

// MonitorFile is a wrapper around g_file_monitor_file().
func (file File) MonitorFile(flags FileMonitorFlags, cancellable Cancellable) (FileMonitor, error) {
	var err glib.Error
	ret0 := C.g_file_monitor_file(file.native(), C.GFileMonitorFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileMonitor{}, err.GoValue()
	}
	return wrapFileMonitor(ret0), nil
}

// MountEnclosingVolumeFinish is a wrapper around g_file_mount_enclosing_volume_finish().
func (location File) MountEnclosingVolumeFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_mount_enclosing_volume_finish(location.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MountMountableFinish is a wrapper around g_file_mount_mountable_finish().
func (file File) MountMountableFinish(result AsyncResult) (File, error) {
	var err glib.Error
	ret0 := C.g_file_mount_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// OpenReadwrite is a wrapper around g_file_open_readwrite().
func (file File) OpenReadwrite(cancellable Cancellable) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_open_readwrite(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// OpenReadwriteFinish is a wrapper around g_file_open_readwrite_finish().
func (file File) OpenReadwriteFinish(res AsyncResult) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_open_readwrite_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// PollMountableFinish is a wrapper around g_file_poll_mountable_finish().
func (file File) PollMountableFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_poll_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// QueryDefaultHandler is a wrapper around g_file_query_default_handler().
func (file File) QueryDefaultHandler(cancellable Cancellable) (AppInfo, error) {
	var err glib.Error
	ret0 := C.g_file_query_default_handler(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return AppInfo{}, err.GoValue()
	}
	return wrapAppInfo(ret0), nil
}

// QueryExists is a wrapper around g_file_query_exists().
func (file File) QueryExists(cancellable Cancellable) bool {
	ret0 := C.g_file_query_exists(file.native(), cancellable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// QueryFileType is a wrapper around g_file_query_file_type().
func (file File) QueryFileType(flags FileQueryInfoFlags, cancellable Cancellable) FileType {
	ret0 := C.g_file_query_file_type(file.native(), C.GFileQueryInfoFlags(flags), cancellable.native())
	return FileType(ret0)
}

// QueryFilesystemInfo is a wrapper around g_file_query_filesystem_info().
func (file File) QueryFilesystemInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_query_filesystem_info(file.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryFilesystemInfoFinish is a wrapper around g_file_query_filesystem_info_finish().
func (file File) QueryFilesystemInfoFinish(res AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_query_filesystem_info_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfo is a wrapper around g_file_query_info().
func (file File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_query_info(file.native(), attributes0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoFinish is a wrapper around g_file_query_info_finish().
func (file File) QueryInfoFinish(res AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_query_info_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryWritableNamespaces is a wrapper around g_file_query_writable_namespaces().
func (file File) QueryWritableNamespaces(cancellable Cancellable) (FileAttributeInfoList, error) {
	var err glib.Error
	ret0 := C.g_file_query_writable_namespaces(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileAttributeInfoList{}, err.GoValue()
	}
	return wrapFileAttributeInfoList(ret0), nil
}

// Read is a wrapper around g_file_read().
func (file File) Read(cancellable Cancellable) (FileInputStream, error) {
	var err glib.Error
	ret0 := C.g_file_read(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInputStream{}, err.GoValue()
	}
	return wrapFileInputStream(ret0), nil
}

// ReadFinish is a wrapper around g_file_read_finish().
func (file File) ReadFinish(res AsyncResult) (FileInputStream, error) {
	var err glib.Error
	ret0 := C.g_file_read_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInputStream{}, err.GoValue()
	}
	return wrapFileInputStream(ret0), nil
}

// Replace is a wrapper around g_file_replace().
func (file File) Replace(etag string, make_backup bool, flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {
	etag0 := C.CString(etag)
	var err glib.Error
	ret0 := C.g_file_replace(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// ReplaceFinish is a wrapper around g_file_replace_finish().
func (file File) ReplaceFinish(res AsyncResult) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_replace_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// ReplaceReadwrite is a wrapper around g_file_replace_readwrite().
func (file File) ReplaceReadwrite(etag string, make_backup bool, flags FileCreateFlags, cancellable Cancellable) (FileIOStream, error) {
	etag0 := C.CString(etag)
	var err glib.Error
	ret0 := C.g_file_replace_readwrite(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// ReplaceReadwriteFinish is a wrapper around g_file_replace_readwrite_finish().
func (file File) ReplaceReadwriteFinish(res AsyncResult) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_replace_readwrite_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// ResolveRelativePath is a wrapper around g_file_resolve_relative_path().
func (file File) ResolveRelativePath(relative_path string) File {
	relative_path0 := C.CString(relative_path)
	ret0 := C.g_file_resolve_relative_path(file.native(), relative_path0)
	C.free(unsafe.Pointer(relative_path0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// SetAttributeByteString is a wrapper around g_file_set_attribute_byte_string().
func (file File) SetAttributeByteString(attribute string, value string, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	value0 := C.CString(value)
	var err glib.Error
	ret0 := C.g_file_set_attribute_byte_string(file.native(), attribute0, value0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))     /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeInt32 is a wrapper around g_file_set_attribute_int32().
func (file File) SetAttributeInt32(attribute string, value int32, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_int32(file.native(), attribute0, C.gint32(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeInt64 is a wrapper around g_file_set_attribute_int64().
func (file File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_int64(file.native(), attribute0, C.gint64(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeString is a wrapper around g_file_set_attribute_string().
func (file File) SetAttributeString(attribute string, value string, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	value0 := C.CString(value)
	var err glib.Error
	ret0 := C.g_file_set_attribute_string(file.native(), attribute0, value0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))     /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeUint32 is a wrapper around g_file_set_attribute_uint32().
func (file File) SetAttributeUint32(attribute string, value uint32, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_uint32(file.native(), attribute0, C.guint32(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeUint64 is a wrapper around g_file_set_attribute_uint64().
func (file File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_uint64(file.native(), attribute0, C.guint64(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributesFinish is a wrapper around g_file_set_attributes_finish().
func (file File) SetAttributesFinish(result AsyncResult) (bool, FileInfo, error) {
	var info0 *C.GFileInfo
	var err glib.Error
	ret0 := C.g_file_set_attributes_finish(file.native(), result.native(), &info0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, FileInfo{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapFileInfo(info0), nil
}

// SetAttributesFromInfo is a wrapper around g_file_set_attributes_from_info().
func (file File) SetAttributesFromInfo(info FileInfo, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_set_attributes_from_info(file.native(), info.native(), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetDisplayName is a wrapper around g_file_set_display_name().
func (file File) SetDisplayName(display_name string, cancellable Cancellable) (File, error) {
	display_name0 := C.CString(display_name)
	var err glib.Error
	ret0 := C.g_file_set_display_name(file.native(), display_name0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// SetDisplayNameFinish is a wrapper around g_file_set_display_name_finish().
func (file File) SetDisplayNameFinish(res AsyncResult) (File, error) {
	var err glib.Error
	ret0 := C.g_file_set_display_name_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// StartMountableFinish is a wrapper around g_file_start_mountable_finish().
func (file File) StartMountableFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_start_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// StopMountableFinish is a wrapper around g_file_stop_mountable_finish().
func (file File) StopMountableFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_stop_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SupportsThreadContexts is a wrapper around g_file_supports_thread_contexts().
func (file File) SupportsThreadContexts() bool {
	ret0 := C.g_file_supports_thread_contexts(file.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Trash is a wrapper around g_file_trash().
func (file File) Trash(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_trash(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// TrashFinish is a wrapper around g_file_trash_finish().
func (file File) TrashFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_trash_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// UnmountMountableWithOperationFinish is a wrapper around g_file_unmount_mountable_with_operation_finish().
func (file File) UnmountMountableWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_unmount_mountable_with_operation_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object Cancellable
type Cancellable struct {
	gobject.Object
}

func (v Cancellable) native() *C.GCancellable {
	return (*C.GCancellable)(v.Ptr)
}
func wrapCancellable(p *C.GCancellable) (v Cancellable) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCancellable(p unsafe.Pointer) (v Cancellable) {
	v.Ptr = p
	return
}

// Object OutputStream
type OutputStream struct {
	gobject.Object
}

func (v OutputStream) native() *C.GOutputStream {
	return (*C.GOutputStream)(v.Ptr)
}
func wrapOutputStream(p *C.GOutputStream) (v OutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapOutputStream(p unsafe.Pointer) (v OutputStream) {
	v.Ptr = p
	return
}

// Object FileOutputStream
type FileOutputStream struct {
	OutputStream
}

func (v FileOutputStream) native() *C.GFileOutputStream {
	return (*C.GFileOutputStream)(v.Ptr)
}
func wrapFileOutputStream(p *C.GFileOutputStream) (v FileOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileOutputStream(p unsafe.Pointer) (v FileOutputStream) {
	v.Ptr = p
	return
}
func (v FileOutputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// QueryInfo is a wrapper around g_file_output_stream_query_info().
func (stream FileOutputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_output_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// Interface Seekable
type Seekable struct {
	Ptr unsafe.Pointer
}

func (v Seekable) native() *C.GSeekable {
	return (*C.GSeekable)(v.Ptr)
}
func wrapSeekable(p *C.GSeekable) Seekable {
	return Seekable{unsafe.Pointer(p)}
}
func WrapSeekable(p unsafe.Pointer) Seekable {
	return Seekable{p}
}

// Object FileInfo
type FileInfo struct {
	gobject.Object
}

func (v FileInfo) native() *C.GFileInfo {
	return (*C.GFileInfo)(v.Ptr)
}
func wrapFileInfo(p *C.GFileInfo) (v FileInfo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileInfo(p unsafe.Pointer) (v FileInfo) {
	v.Ptr = p
	return
}

// Object Application
type Application struct {
	gobject.Object
}

func (v Application) native() *C.GApplication {
	return (*C.GApplication)(v.Ptr)
}
func wrapApplication(p *C.GApplication) (v Application) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapApplication(p unsafe.Pointer) (v Application) {
	v.Ptr = p
	return
}
func (v Application) ActionGroup() ActionGroup {
	return WrapActionGroup(v.Ptr)
}
func (v Application) ActionMap() ActionMap {
	return WrapActionMap(v.Ptr)
}

// ApplicationNew is a wrapper around g_application_new().
func ApplicationNew(application_id string, flags ApplicationFlags) Application {
	application_id0 := (*C.gchar)(C.CString(application_id))
	ret0 := C.g_application_new(application_id0, C.GApplicationFlags(flags))
	C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/
	return wrapApplication(ret0)
}

// Activate is a wrapper around g_application_activate().
func (application Application) Activate() {
	C.g_application_activate(application.native())
}

// BindBusyProperty is a wrapper around g_application_bind_busy_property().
func (application Application) BindBusyProperty(object gobject.Object, property string) {
	property0 := (*C.gchar)(C.CString(property))
	C.g_application_bind_busy_property(application.native(), C.gpointer((C.gpointer)(object.Ptr)), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// GetApplicationId is a wrapper around g_application_get_application_id().
func (application Application) GetApplicationId() string {
	ret0 := C.g_application_get_application_id(application.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDbusObjectPath is a wrapper around g_application_get_dbus_object_path().
func (application Application) GetDbusObjectPath() string {
	ret0 := C.g_application_get_dbus_object_path(application.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetFlags is a wrapper around g_application_get_flags().
func (application Application) GetFlags() ApplicationFlags {
	ret0 := C.g_application_get_flags(application.native())
	return ApplicationFlags(ret0)
}

// GetInactivityTimeout is a wrapper around g_application_get_inactivity_timeout().
func (application Application) GetInactivityTimeout() uint {
	ret0 := C.g_application_get_inactivity_timeout(application.native())
	return uint(ret0)
}

// GetIsBusy is a wrapper around g_application_get_is_busy().
func (application Application) GetIsBusy() bool {
	ret0 := C.g_application_get_is_busy(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsRegistered is a wrapper around g_application_get_is_registered().
func (application Application) GetIsRegistered() bool {
	ret0 := C.g_application_get_is_registered(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsRemote is a wrapper around g_application_get_is_remote().
func (application Application) GetIsRemote() bool {
	ret0 := C.g_application_get_is_remote(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetResourceBasePath is a wrapper around g_application_get_resource_base_path().
func (application Application) GetResourceBasePath() string {
	ret0 := C.g_application_get_resource_base_path(application.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Hold is a wrapper around g_application_hold().
func (application Application) Hold() {
	C.g_application_hold(application.native())
}

// MarkBusy is a wrapper around g_application_mark_busy().
func (application Application) MarkBusy() {
	C.g_application_mark_busy(application.native())
}

// Open is a wrapper around g_application_open().
func (application Application) Open(files []File, hint string) {
	files0 := make([]*C.GFile, len(files))
	for idx, elemG := range files {
		files0[idx] = elemG.native()
	}
	var files0Ptr **C.GFile
	if len(files0) > 0 {
		files0Ptr = &files0[0]
	}
	hint0 := (*C.gchar)(C.CString(hint))
	C.g_application_open(application.native(), files0Ptr, C.gint(len(files)), hint0)
	C.free(unsafe.Pointer(hint0)) /*ch:<stdlib.h>*/
}

// Quit is a wrapper around g_application_quit().
func (application Application) Quit() {
	C.g_application_quit(application.native())
}

// Register is a wrapper around g_application_register().
func (application Application) Register(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_application_register(application.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Release is a wrapper around g_application_release().
func (application Application) Release() {
	C.g_application_release(application.native())
}

// Run is a wrapper around g_application_run().
func (application Application) Run(argv []string) int {
	argv0 := make([]*C.char, len(argv))
	for idx, elemG := range argv {
		elem := C.CString(elemG)
		argv0[idx] = elem
	}
	var argv0Ptr **C.char
	if len(argv0) > 0 {
		argv0Ptr = &argv0[0]
	}
	ret0 := C.g_application_run(application.native(), C.int(len(argv)), argv0Ptr)
	for _, elem := range argv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return int(ret0)
}

// SetApplicationId is a wrapper around g_application_set_application_id().
func (application Application) SetApplicationId(application_id string) {
	application_id0 := (*C.gchar)(C.CString(application_id))
	C.g_application_set_application_id(application.native(), application_id0)
	C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/
}

// SetDefault is a wrapper around g_application_set_default().
func (application Application) SetDefault() {
	C.g_application_set_default(application.native())
}

// SetFlags is a wrapper around g_application_set_flags().
func (application Application) SetFlags(flags ApplicationFlags) {
	C.g_application_set_flags(application.native(), C.GApplicationFlags(flags))
}

// SetInactivityTimeout is a wrapper around g_application_set_inactivity_timeout().
func (application Application) SetInactivityTimeout(inactivity_timeout uint) {
	C.g_application_set_inactivity_timeout(application.native(), C.guint(inactivity_timeout))
}

// SetResourceBasePath is a wrapper around g_application_set_resource_base_path().
func (application Application) SetResourceBasePath(resource_path string) {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	C.g_application_set_resource_base_path(application.native(), resource_path0)
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
}

// UnbindBusyProperty is a wrapper around g_application_unbind_busy_property().
func (application Application) UnbindBusyProperty(object gobject.Object, property string) {
	property0 := (*C.gchar)(C.CString(property))
	C.g_application_unbind_busy_property(application.native(), C.gpointer((C.gpointer)(object.Ptr)), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// UnmarkBusy is a wrapper around g_application_unmark_busy().
func (application Application) UnmarkBusy() {
	C.g_application_unmark_busy(application.native())
}

// WithdrawNotification is a wrapper around g_application_withdraw_notification().
func (application Application) WithdrawNotification(id string) {
	id0 := (*C.gchar)(C.CString(id))
	C.g_application_withdraw_notification(application.native(), id0)
	C.free(unsafe.Pointer(id0)) /*ch:<stdlib.h>*/
}

// ApplicationGetDefault is a wrapper around g_application_get_default().
func ApplicationGetDefault() Application {
	ret0 := C.g_application_get_default()
	return wrapApplication(ret0)
}

// ApplicationIdIsValid is a wrapper around g_application_id_is_valid().
func ApplicationIdIsValid(application_id string) bool {
	application_id0 := (*C.gchar)(C.CString(application_id))
	ret0 := C.g_application_id_is_valid(application_id0)
	C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))         /*go:.util*/
}

// Interface ActionMap
type ActionMap struct {
	Ptr unsafe.Pointer
}

func (v ActionMap) native() *C.GActionMap {
	return (*C.GActionMap)(v.Ptr)
}
func wrapActionMap(p *C.GActionMap) ActionMap {
	return ActionMap{unsafe.Pointer(p)}
}
func WrapActionMap(p unsafe.Pointer) ActionMap {
	return ActionMap{p}
}

// AddAction is a wrapper around g_action_map_add_action().
func (action_map ActionMap) AddAction(action Action) {
	C.g_action_map_add_action(action_map.native(), action.native())
}

// LookupAction is a wrapper around g_action_map_lookup_action().
func (action_map ActionMap) LookupAction(action_name string) Action {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_map_lookup_action(action_map.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return wrapAction(ret0)
}

// RemoveAction is a wrapper around g_action_map_remove_action().
func (action_map ActionMap) RemoveAction(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_map_remove_action(action_map.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// Interface ActionGroup
type ActionGroup struct {
	Ptr unsafe.Pointer
}

func (v ActionGroup) native() *C.GActionGroup {
	return (*C.GActionGroup)(v.Ptr)
}
func wrapActionGroup(p *C.GActionGroup) ActionGroup {
	return ActionGroup{unsafe.Pointer(p)}
}
func WrapActionGroup(p unsafe.Pointer) ActionGroup {
	return ActionGroup{p}
}

// ActionAdded is a wrapper around g_action_group_action_added().
func (action_group ActionGroup) ActionAdded(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_added(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActionEnabledChanged is a wrapper around g_action_group_action_enabled_changed().
func (action_group ActionGroup) ActionEnabledChanged(action_name string, enabled bool) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_enabled_changed(action_group.native(), action_name0, C.gboolean(util.Bool2Int(enabled)) /*go:.util*/)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActionRemoved is a wrapper around g_action_group_action_removed().
func (action_group ActionGroup) ActionRemoved(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_removed(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActionStateChanged is a wrapper around g_action_group_action_state_changed().
func (action_group ActionGroup) ActionStateChanged(action_name string, state glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_state_changed(action_group.native(), action_name0, (*C.GVariant)(state.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActivateAction is a wrapper around g_action_group_activate_action().
func (action_group ActionGroup) ActivateAction(action_name string, parameter glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_activate_action(action_group.native(), action_name0, (*C.GVariant)(parameter.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ChangeActionState is a wrapper around g_action_group_change_action_state().
func (action_group ActionGroup) ChangeActionState(action_name string, value glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_change_action_state(action_group.native(), action_name0, (*C.GVariant)(value.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// GetActionEnabled is a wrapper around g_action_group_get_action_enabled().
func (action_group ActionGroup) GetActionEnabled(action_name string) bool {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_enabled(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// GetActionParameterType is a wrapper around g_action_group_get_action_parameter_type().
func (action_group ActionGroup) GetActionParameterType(action_name string) glib.VariantType {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_parameter_type(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))              /*ch:<stdlib.h>*/
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetActionState is a wrapper around g_action_group_get_action_state().
func (action_group ActionGroup) GetActionState(action_name string) glib.Variant {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_state(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))          /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetActionStateHint is a wrapper around g_action_group_get_action_state_hint().
func (action_group ActionGroup) GetActionStateHint(action_name string) glib.Variant {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_state_hint(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))          /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetActionStateType is a wrapper around g_action_group_get_action_state_type().
func (action_group ActionGroup) GetActionStateType(action_name string) glib.VariantType {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_state_type(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))              /*ch:<stdlib.h>*/
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// HasAction is a wrapper around g_action_group_has_action().
func (action_group ActionGroup) HasAction(action_name string) bool {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_has_action(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// ListActions is a wrapper around g_action_group_list_actions().
func (action_group ActionGroup) ListActions() []string {
	ret0 := C.g_action_group_list_actions(action_group.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// QueryAction is a wrapper around g_action_group_query_action().
func (action_group ActionGroup) QueryAction(action_name string) (bool, bool, glib.VariantType, glib.VariantType, glib.Variant, glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	var enabled0 C.gboolean
	var parameter_type0 *C.GVariantType
	var state_type0 *C.GVariantType
	var state_hint0 *C.GVariant
	var state0 *C.GVariant
	ret0 := C.g_action_group_query_action(action_group.native(), action_name0, &enabled0, &parameter_type0, &state_type0, &state_hint0, &state0)
	C.free(unsafe.Pointer(action_name0))                                                                                                                                                                                                                                                                                              /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/, util.Int2Bool(int(enabled0)) /*go:.util*/, glib.WrapVariantType(unsafe.Pointer(parameter_type0)) /*gir:GLib*/, glib.WrapVariantType(unsafe.Pointer(state_type0)) /*gir:GLib*/, glib.WrapVariant(unsafe.Pointer(state_hint0)) /*gir:GLib*/, glib.WrapVariant(unsafe.Pointer(state0)) /*gir:GLib*/
}

// Object SimpleActionGroup
type SimpleActionGroup struct {
	gobject.Object
}

func (v SimpleActionGroup) native() *C.GSimpleActionGroup {
	return (*C.GSimpleActionGroup)(v.Ptr)
}
func wrapSimpleActionGroup(p *C.GSimpleActionGroup) (v SimpleActionGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimpleActionGroup(p unsafe.Pointer) (v SimpleActionGroup) {
	v.Ptr = p
	return
}
func (v SimpleActionGroup) ActionGroup() ActionGroup {
	return WrapActionGroup(v.Ptr)
}
func (v SimpleActionGroup) ActionMap() ActionMap {
	return WrapActionMap(v.Ptr)
}

// SimpleActionGroupNew is a wrapper around g_simple_action_group_new().
func SimpleActionGroupNew() SimpleActionGroup {
	ret0 := C.g_simple_action_group_new()
	return wrapSimpleActionGroup(ret0)
}

// Object Notification
type Notification struct {
	gobject.Object
}

func (v Notification) native() *C.GNotification {
	return (*C.GNotification)(v.Ptr)
}
func wrapNotification(p *C.GNotification) (v Notification) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNotification(p unsafe.Pointer) (v Notification) {
	v.Ptr = p
	return
}

// NotificationNew is a wrapper around g_notification_new().
func NotificationNew(title string) Notification {
	title0 := (*C.gchar)(C.CString(title))
	ret0 := C.g_notification_new(title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
	return wrapNotification(ret0)
}

// AddButton is a wrapper around g_notification_add_button().
func (notification Notification) AddButton(label string, detailed_action string) {
	label0 := (*C.gchar)(C.CString(label))
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_notification_add_button(notification.native(), label0, detailed_action0)
	C.free(unsafe.Pointer(label0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// AddButtonWithTargetValue is a wrapper around g_notification_add_button_with_target_value().
func (notification Notification) AddButtonWithTargetValue(label string, action string, target glib.Variant) {
	label0 := (*C.gchar)(C.CString(label))
	action0 := (*C.gchar)(C.CString(action))
	C.g_notification_add_button_with_target_value(notification.native(), label0, action0, (*C.GVariant)(target.Ptr))
	C.free(unsafe.Pointer(label0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(action0)) /*ch:<stdlib.h>*/
}

// SetBody is a wrapper around g_notification_set_body().
func (notification Notification) SetBody(body string) {
	body0 := (*C.gchar)(C.CString(body))
	C.g_notification_set_body(notification.native(), body0)
	C.free(unsafe.Pointer(body0)) /*ch:<stdlib.h>*/
}

// SetDefaultAction is a wrapper around g_notification_set_default_action().
func (notification Notification) SetDefaultAction(detailed_action string) {
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_notification_set_default_action(notification.native(), detailed_action0)
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// SetDefaultActionAndTargetValue is a wrapper around g_notification_set_default_action_and_target_value().
func (notification Notification) SetDefaultActionAndTargetValue(action string, target glib.Variant) {
	action0 := (*C.gchar)(C.CString(action))
	C.g_notification_set_default_action_and_target_value(notification.native(), action0, (*C.GVariant)(target.Ptr))
	C.free(unsafe.Pointer(action0)) /*ch:<stdlib.h>*/
}

// SetIcon is a wrapper around g_notification_set_icon().
func (notification Notification) SetIcon(icon Icon) {
	C.g_notification_set_icon(notification.native(), icon.native())
}

// SetPriority is a wrapper around g_notification_set_priority().
func (notification Notification) SetPriority(priority NotificationPriority) {
	C.g_notification_set_priority(notification.native(), C.GNotificationPriority(priority))
}

// SetTitle is a wrapper around g_notification_set_title().
func (notification Notification) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.g_notification_set_title(notification.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// Interface Icon
type Icon struct {
	Ptr unsafe.Pointer
}

func (v Icon) native() *C.GIcon {
	return (*C.GIcon)(v.Ptr)
}
func wrapIcon(p *C.GIcon) Icon {
	return Icon{unsafe.Pointer(p)}
}
func WrapIcon(p unsafe.Pointer) Icon {
	return Icon{p}
}

// Equal is a wrapper around g_icon_equal().
func (icon1 Icon) Equal(icon2 Icon) bool {
	ret0 := C.g_icon_equal(icon1.native(), icon2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Serialize is a wrapper around g_icon_serialize().
func (icon Icon) Serialize() glib.Variant {
	ret0 := C.g_icon_serialize(icon.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// ToString is a wrapper around g_icon_to_string().
func (icon Icon) ToString() string {
	ret0 := C.g_icon_to_string(icon.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Interface AsyncResult
type AsyncResult struct {
	Ptr unsafe.Pointer
}

func (v AsyncResult) native() *C.GAsyncResult {
	return (*C.GAsyncResult)(v.Ptr)
}
func wrapAsyncResult(p *C.GAsyncResult) AsyncResult {
	return AsyncResult{unsafe.Pointer(p)}
}
func WrapAsyncResult(p unsafe.Pointer) AsyncResult {
	return AsyncResult{p}
}

// GetSourceObject is a wrapper around g_async_result_get_source_object().
func (res AsyncResult) GetSourceObject() gobject.Object {
	ret0 := C.g_async_result_get_source_object(res.native())
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetUserData is a wrapper around g_async_result_get_user_data().
func (res AsyncResult) GetUserData() unsafe.Pointer {
	ret0 := C.g_async_result_get_user_data(res.native())
	return unsafe.Pointer(ret0)
}

// IsTagged is a wrapper around g_async_result_is_tagged().
func (res AsyncResult) IsTagged(source_tag unsafe.Pointer) bool {
	ret0 := C.g_async_result_is_tagged(res.native(), C.gpointer(source_tag))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LegacyPropagateError is a wrapper around g_async_result_legacy_propagate_error().
func (res AsyncResult) LegacyPropagateError() (bool, error) {
	var err glib.Error
	ret0 := C.g_async_result_legacy_propagate_error(res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object IOStream
type IOStream struct {
	gobject.Object
}

func (v IOStream) native() *C.GIOStream {
	return (*C.GIOStream)(v.Ptr)
}
func wrapIOStream(p *C.GIOStream) (v IOStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIOStream(p unsafe.Pointer) (v IOStream) {
	v.Ptr = p
	return
}

// ClearPending is a wrapper around g_io_stream_clear_pending().
func (stream IOStream) ClearPending() {
	C.g_io_stream_clear_pending(stream.native())
}

// Close is a wrapper around g_io_stream_close().
func (stream IOStream) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_close(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseFinish is a wrapper around g_io_stream_close_finish().
func (stream IOStream) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_close_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetInputStream is a wrapper around g_io_stream_get_input_stream().
func (stream IOStream) GetInputStream() InputStream {
	ret0 := C.g_io_stream_get_input_stream(stream.native())
	return wrapInputStream(ret0)
}

// GetOutputStream is a wrapper around g_io_stream_get_output_stream().
func (stream IOStream) GetOutputStream() OutputStream {
	ret0 := C.g_io_stream_get_output_stream(stream.native())
	return wrapOutputStream(ret0)
}

// HasPending is a wrapper around g_io_stream_has_pending().
func (stream IOStream) HasPending() bool {
	ret0 := C.g_io_stream_has_pending(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_io_stream_is_closed().
func (stream IOStream) IsClosed() bool {
	ret0 := C.g_io_stream_is_closed(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetPending is a wrapper around g_io_stream_set_pending().
func (stream IOStream) SetPending() (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_set_pending(stream.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// IOStreamSpliceFinish is a wrapper around g_io_stream_splice_finish().
func IOStreamSpliceFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_splice_finish(result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object InputStream
type InputStream struct {
	gobject.Object
}

func (v InputStream) native() *C.GInputStream {
	return (*C.GInputStream)(v.Ptr)
}
func wrapInputStream(p *C.GInputStream) (v InputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInputStream(p unsafe.Pointer) (v InputStream) {
	v.Ptr = p
	return
}

// ClearPending is a wrapper around g_input_stream_clear_pending().
func (stream InputStream) ClearPending() {
	C.g_input_stream_clear_pending(stream.native())
}

// Close is a wrapper around g_input_stream_close().
func (stream InputStream) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_input_stream_close(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseFinish is a wrapper around g_input_stream_close_finish().
func (stream InputStream) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_input_stream_close_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// HasPending is a wrapper around g_input_stream_has_pending().
func (stream InputStream) HasPending() bool {
	ret0 := C.g_input_stream_has_pending(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_input_stream_is_closed().
func (stream InputStream) IsClosed() bool {
	ret0 := C.g_input_stream_is_closed(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ReadAllFinish is a wrapper around g_input_stream_read_all_finish().
func (stream InputStream) ReadAllFinish(result AsyncResult) (bool, uint, error) {
	var bytes_read0 C.gsize
	var err glib.Error
	ret0 := C.g_input_stream_read_all_finish(stream.native(), result.native(), &bytes_read0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint(bytes_read0), nil
}

// ReadBytes is a wrapper around g_input_stream_read_bytes().
func (stream InputStream) ReadBytes(count uint, cancellable Cancellable) (glib.Bytes, error) {
	var err glib.Error
	ret0 := C.g_input_stream_read_bytes(stream.native(), C.gsize(count), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// ReadBytesFinish is a wrapper around g_input_stream_read_bytes_finish().
func (stream InputStream) ReadBytesFinish(result AsyncResult) (glib.Bytes, error) {
	var err glib.Error
	ret0 := C.g_input_stream_read_bytes_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// ReadFinish is a wrapper around g_input_stream_read_finish().
func (stream InputStream) ReadFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_input_stream_read_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SetPending is a wrapper around g_input_stream_set_pending().
func (stream InputStream) SetPending() (bool, error) {
	var err glib.Error
	ret0 := C.g_input_stream_set_pending(stream.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Skip is a wrapper around g_input_stream_skip().
func (stream InputStream) Skip(count uint, cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_input_stream_skip(stream.native(), C.gsize(count), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SkipFinish is a wrapper around g_input_stream_skip_finish().
func (stream InputStream) SkipFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_input_stream_skip_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// Object FileIOStream
type FileIOStream struct {
	IOStream
}

func (v FileIOStream) native() *C.GFileIOStream {
	return (*C.GFileIOStream)(v.Ptr)
}
func wrapFileIOStream(p *C.GFileIOStream) (v FileIOStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileIOStream(p unsafe.Pointer) (v FileIOStream) {
	v.Ptr = p
	return
}
func (v FileIOStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// GetEtag is a wrapper around g_file_io_stream_get_etag().
func (stream FileIOStream) GetEtag() string {
	ret0 := C.g_file_io_stream_get_etag(stream.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// QueryInfo is a wrapper around g_file_io_stream_query_info().
func (stream FileIOStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_io_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoFinish is a wrapper around g_file_io_stream_query_info_finish().
func (stream FileIOStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_io_stream_query_info_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// Object FileEnumerator
type FileEnumerator struct {
	gobject.Object
}

func (v FileEnumerator) native() *C.GFileEnumerator {
	return (*C.GFileEnumerator)(v.Ptr)
}
func wrapFileEnumerator(p *C.GFileEnumerator) (v FileEnumerator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileEnumerator(p unsafe.Pointer) (v FileEnumerator) {
	v.Ptr = p
	return
}

// Close is a wrapper around g_file_enumerator_close().
func (enumerator FileEnumerator) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_close(enumerator.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseFinish is a wrapper around g_file_enumerator_close_finish().
func (enumerator FileEnumerator) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_close_finish(enumerator.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetChild is a wrapper around g_file_enumerator_get_child().
func (enumerator FileEnumerator) GetChild(info FileInfo) File {
	ret0 := C.g_file_enumerator_get_child(enumerator.native(), info.native())
	return wrapFile(ret0)
}

// GetContainer is a wrapper around g_file_enumerator_get_container().
func (enumerator FileEnumerator) GetContainer() File {
	ret0 := C.g_file_enumerator_get_container(enumerator.native())
	return wrapFile(ret0)
}

// HasPending is a wrapper around g_file_enumerator_has_pending().
func (enumerator FileEnumerator) HasPending() bool {
	ret0 := C.g_file_enumerator_has_pending(enumerator.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_file_enumerator_is_closed().
func (enumerator FileEnumerator) IsClosed() bool {
	ret0 := C.g_file_enumerator_is_closed(enumerator.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Iterate is a wrapper around g_file_enumerator_iterate().
func (direnum FileEnumerator) Iterate(cancellable Cancellable) (bool, FileInfo, File, error) {
	var out_info0 *C.GFileInfo
	var out_child0 *C.GFile
	var err glib.Error
	ret0 := C.g_file_enumerator_iterate(direnum.native(), &out_info0, &out_child0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, FileInfo{}, File{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapFileInfo(out_info0), wrapFile(out_child0), nil
}

// NextFile is a wrapper around g_file_enumerator_next_file().
func (enumerator FileEnumerator) NextFile(cancellable Cancellable) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_next_file(enumerator.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// SetPending is a wrapper around g_file_enumerator_set_pending().
func (enumerator FileEnumerator) SetPending(pending bool) {
	C.g_file_enumerator_set_pending(enumerator.native(), C.gboolean(util.Bool2Int(pending)) /*go:.util*/)
}

// Interface Mount
type Mount struct {
	Ptr unsafe.Pointer
}

func (v Mount) native() *C.GMount {
	return (*C.GMount)(v.Ptr)
}
func wrapMount(p *C.GMount) Mount {
	return Mount{unsafe.Pointer(p)}
}
func WrapMount(p unsafe.Pointer) Mount {
	return Mount{p}
}

// CanEject is a wrapper around g_mount_can_eject().
func (mount Mount) CanEject() bool {
	ret0 := C.g_mount_can_eject(mount.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanUnmount is a wrapper around g_mount_can_unmount().
func (mount Mount) CanUnmount() bool {
	ret0 := C.g_mount_can_unmount(mount.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EjectWithOperationFinish is a wrapper around g_mount_eject_with_operation_finish().
func (mount Mount) EjectWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_mount_eject_with_operation_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetDefaultLocation is a wrapper around g_mount_get_default_location().
func (mount Mount) GetDefaultLocation() File {
	ret0 := C.g_mount_get_default_location(mount.native())
	return wrapFile(ret0)
}

// GetDrive is a wrapper around g_mount_get_drive().
func (mount Mount) GetDrive() Drive {
	ret0 := C.g_mount_get_drive(mount.native())
	return wrapDrive(ret0)
}

// GetIcon is a wrapper around g_mount_get_icon().
func (mount Mount) GetIcon() Icon {
	ret0 := C.g_mount_get_icon(mount.native())
	return wrapIcon(ret0)
}

// GetName is a wrapper around g_mount_get_name().
func (mount Mount) GetName() string {
	ret0 := C.g_mount_get_name(mount.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetRoot is a wrapper around g_mount_get_root().
func (mount Mount) GetRoot() File {
	ret0 := C.g_mount_get_root(mount.native())
	return wrapFile(ret0)
}

// GetSortKey is a wrapper around g_mount_get_sort_key().
func (mount Mount) GetSortKey() string {
	ret0 := C.g_mount_get_sort_key(mount.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSymbolicIcon is a wrapper around g_mount_get_symbolic_icon().
func (mount Mount) GetSymbolicIcon() Icon {
	ret0 := C.g_mount_get_symbolic_icon(mount.native())
	return wrapIcon(ret0)
}

// GetUuid is a wrapper around g_mount_get_uuid().
func (mount Mount) GetUuid() string {
	ret0 := C.g_mount_get_uuid(mount.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetVolume is a wrapper around g_mount_get_volume().
func (mount Mount) GetVolume() Volume {
	ret0 := C.g_mount_get_volume(mount.native())
	return wrapVolume(ret0)
}

// GuessContentTypeFinish is a wrapper around g_mount_guess_content_type_finish().
func (mount Mount) GuessContentTypeFinish(result AsyncResult) ([]string, error) {
	var err glib.Error
	ret0 := C.g_mount_guess_content_type_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GuessContentTypeSync is a wrapper around g_mount_guess_content_type_sync().
func (mount Mount) GuessContentTypeSync(force_rescan bool, cancellable Cancellable) ([]string, error) {
	var err glib.Error
	ret0 := C.g_mount_guess_content_type_sync(mount.native(), C.gboolean(util.Bool2Int(force_rescan)) /*go:.util*/, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// IsShadowed is a wrapper around g_mount_is_shadowed().
func (mount Mount) IsShadowed() bool {
	ret0 := C.g_mount_is_shadowed(mount.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RemountFinish is a wrapper around g_mount_remount_finish().
func (mount Mount) RemountFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_mount_remount_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Shadow is a wrapper around g_mount_shadow().
func (mount Mount) Shadow() {
	C.g_mount_shadow(mount.native())
}

// UnmountWithOperationFinish is a wrapper around g_mount_unmount_with_operation_finish().
func (mount Mount) UnmountWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_mount_unmount_with_operation_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Unshadow is a wrapper around g_mount_unshadow().
func (mount Mount) Unshadow() {
	C.g_mount_unshadow(mount.native())
}

// Interface Drive
type Drive struct {
	Ptr unsafe.Pointer
}

func (v Drive) native() *C.GDrive {
	return (*C.GDrive)(v.Ptr)
}
func wrapDrive(p *C.GDrive) Drive {
	return Drive{unsafe.Pointer(p)}
}
func WrapDrive(p unsafe.Pointer) Drive {
	return Drive{p}
}

// CanEject is a wrapper around g_drive_can_eject().
func (drive Drive) CanEject() bool {
	ret0 := C.g_drive_can_eject(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanPollForMedia is a wrapper around g_drive_can_poll_for_media().
func (drive Drive) CanPollForMedia() bool {
	ret0 := C.g_drive_can_poll_for_media(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanStart is a wrapper around g_drive_can_start().
func (drive Drive) CanStart() bool {
	ret0 := C.g_drive_can_start(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanStartDegraded is a wrapper around g_drive_can_start_degraded().
func (drive Drive) CanStartDegraded() bool {
	ret0 := C.g_drive_can_start_degraded(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanStop is a wrapper around g_drive_can_stop().
func (drive Drive) CanStop() bool {
	ret0 := C.g_drive_can_stop(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EjectWithOperationFinish is a wrapper around g_drive_eject_with_operation_finish().
func (drive Drive) EjectWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_eject_with_operation_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EnumerateIdentifiers is a wrapper around g_drive_enumerate_identifiers().
func (drive Drive) EnumerateIdentifiers() []string {
	ret0 := C.g_drive_enumerate_identifiers(drive.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetIcon is a wrapper around g_drive_get_icon().
func (drive Drive) GetIcon() Icon {
	ret0 := C.g_drive_get_icon(drive.native())
	return wrapIcon(ret0)
}

// GetIdentifier is a wrapper around g_drive_get_identifier().
func (drive Drive) GetIdentifier(kind string) string {
	kind0 := C.CString(kind)
	ret0 := C.g_drive_get_identifier(drive.native(), kind0)
	C.free(unsafe.Pointer(kind0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetName is a wrapper around g_drive_get_name().
func (drive Drive) GetName() string {
	ret0 := C.g_drive_get_name(drive.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetSortKey is a wrapper around g_drive_get_sort_key().
func (drive Drive) GetSortKey() string {
	ret0 := C.g_drive_get_sort_key(drive.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetStartStopType is a wrapper around g_drive_get_start_stop_type().
func (drive Drive) GetStartStopType() DriveStartStopType {
	ret0 := C.g_drive_get_start_stop_type(drive.native())
	return DriveStartStopType(ret0)
}

// GetSymbolicIcon is a wrapper around g_drive_get_symbolic_icon().
func (drive Drive) GetSymbolicIcon() Icon {
	ret0 := C.g_drive_get_symbolic_icon(drive.native())
	return wrapIcon(ret0)
}

// HasMedia is a wrapper around g_drive_has_media().
func (drive Drive) HasMedia() bool {
	ret0 := C.g_drive_has_media(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasVolumes is a wrapper around g_drive_has_volumes().
func (drive Drive) HasVolumes() bool {
	ret0 := C.g_drive_has_volumes(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsMediaCheckAutomatic is a wrapper around g_drive_is_media_check_automatic().
func (drive Drive) IsMediaCheckAutomatic() bool {
	ret0 := C.g_drive_is_media_check_automatic(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsMediaRemovable is a wrapper around g_drive_is_media_removable().
func (drive Drive) IsMediaRemovable() bool {
	ret0 := C.g_drive_is_media_removable(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsRemovable is a wrapper around g_drive_is_removable().
func (drive Drive) IsRemovable() bool {
	ret0 := C.g_drive_is_removable(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PollForMediaFinish is a wrapper around g_drive_poll_for_media_finish().
func (drive Drive) PollForMediaFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_poll_for_media_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// StartFinish is a wrapper around g_drive_start_finish().
func (drive Drive) StartFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_start_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// StopFinish is a wrapper around g_drive_stop_finish().
func (drive Drive) StopFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_stop_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Interface Volume
type Volume struct {
	Ptr unsafe.Pointer
}

func (v Volume) native() *C.GVolume {
	return (*C.GVolume)(v.Ptr)
}
func wrapVolume(p *C.GVolume) Volume {
	return Volume{unsafe.Pointer(p)}
}
func WrapVolume(p unsafe.Pointer) Volume {
	return Volume{p}
}

// CanEject is a wrapper around g_volume_can_eject().
func (volume Volume) CanEject() bool {
	ret0 := C.g_volume_can_eject(volume.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanMount is a wrapper around g_volume_can_mount().
func (volume Volume) CanMount() bool {
	ret0 := C.g_volume_can_mount(volume.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EjectWithOperationFinish is a wrapper around g_volume_eject_with_operation_finish().
func (volume Volume) EjectWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_volume_eject_with_operation_finish(volume.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EnumerateIdentifiers is a wrapper around g_volume_enumerate_identifiers().
func (volume Volume) EnumerateIdentifiers() []string {
	ret0 := C.g_volume_enumerate_identifiers(volume.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetActivationRoot is a wrapper around g_volume_get_activation_root().
func (volume Volume) GetActivationRoot() File {
	ret0 := C.g_volume_get_activation_root(volume.native())
	return wrapFile(ret0)
}

// GetDrive is a wrapper around g_volume_get_drive().
func (volume Volume) GetDrive() Drive {
	ret0 := C.g_volume_get_drive(volume.native())
	return wrapDrive(ret0)
}

// GetIcon is a wrapper around g_volume_get_icon().
func (volume Volume) GetIcon() Icon {
	ret0 := C.g_volume_get_icon(volume.native())
	return wrapIcon(ret0)
}

// GetIdentifier is a wrapper around g_volume_get_identifier().
func (volume Volume) GetIdentifier(kind string) string {
	kind0 := C.CString(kind)
	ret0 := C.g_volume_get_identifier(volume.native(), kind0)
	C.free(unsafe.Pointer(kind0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetMount is a wrapper around g_volume_get_mount().
func (volume Volume) GetMount() Mount {
	ret0 := C.g_volume_get_mount(volume.native())
	return wrapMount(ret0)
}

// GetName is a wrapper around g_volume_get_name().
func (volume Volume) GetName() string {
	ret0 := C.g_volume_get_name(volume.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetSortKey is a wrapper around g_volume_get_sort_key().
func (volume Volume) GetSortKey() string {
	ret0 := C.g_volume_get_sort_key(volume.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSymbolicIcon is a wrapper around g_volume_get_symbolic_icon().
func (volume Volume) GetSymbolicIcon() Icon {
	ret0 := C.g_volume_get_symbolic_icon(volume.native())
	return wrapIcon(ret0)
}

// GetUuid is a wrapper around g_volume_get_uuid().
func (volume Volume) GetUuid() string {
	ret0 := C.g_volume_get_uuid(volume.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// MountFinish is a wrapper around g_volume_mount_finish().
func (volume Volume) MountFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_volume_mount_finish(volume.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ShouldAutomount is a wrapper around g_volume_should_automount().
func (volume Volume) ShouldAutomount() bool {
	ret0 := C.g_volume_should_automount(volume.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object FileMonitor
type FileMonitor struct {
	gobject.Object
}

func (v FileMonitor) native() *C.GFileMonitor {
	return (*C.GFileMonitor)(v.Ptr)
}
func wrapFileMonitor(p *C.GFileMonitor) (v FileMonitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileMonitor(p unsafe.Pointer) (v FileMonitor) {
	v.Ptr = p
	return
}

// Cancel is a wrapper around g_file_monitor_cancel().
func (monitor FileMonitor) Cancel() bool {
	ret0 := C.g_file_monitor_cancel(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EmitEvent is a wrapper around g_file_monitor_emit_event().
func (monitor FileMonitor) EmitEvent(child File, other_file File, event_type FileMonitorEvent) {
	C.g_file_monitor_emit_event(monitor.native(), child.native(), other_file.native(), C.GFileMonitorEvent(event_type))
}

// IsCancelled is a wrapper around g_file_monitor_is_cancelled().
func (monitor FileMonitor) IsCancelled() bool {
	ret0 := C.g_file_monitor_is_cancelled(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetRateLimit is a wrapper around g_file_monitor_set_rate_limit().
func (monitor FileMonitor) SetRateLimit(limit_msecs int) {
	C.g_file_monitor_set_rate_limit(monitor.native(), C.gint(limit_msecs))
}

// Struct FileAttributeInfoList
type FileAttributeInfoList struct {
	Ptr unsafe.Pointer
}

func (v FileAttributeInfoList) native() *C.GFileAttributeInfoList {
	return (*C.GFileAttributeInfoList)(v.Ptr)
}
func wrapFileAttributeInfoList(p *C.GFileAttributeInfoList) FileAttributeInfoList {
	return FileAttributeInfoList{unsafe.Pointer(p)}
}
func WrapFileAttributeInfoList(p unsafe.Pointer) FileAttributeInfoList {
	return FileAttributeInfoList{p}
}

// FileAttributeInfoListNew is a wrapper around g_file_attribute_info_list_new().
func FileAttributeInfoListNew() FileAttributeInfoList {
	ret0 := C.g_file_attribute_info_list_new()
	return wrapFileAttributeInfoList(ret0)
}

// Dup is a wrapper around g_file_attribute_info_list_dup().
func (list FileAttributeInfoList) Dup() FileAttributeInfoList {
	ret0 := C.g_file_attribute_info_list_dup(list.native())
	return wrapFileAttributeInfoList(ret0)
}

// Lookup is a wrapper around g_file_attribute_info_list_lookup().
func (list FileAttributeInfoList) Lookup(name string) FileAttributeInfo {
	name0 := C.CString(name)
	ret0 := C.g_file_attribute_info_list_lookup(list.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapFileAttributeInfo(ret0)
}

// Ref is a wrapper around g_file_attribute_info_list_ref().
func (list FileAttributeInfoList) Ref() FileAttributeInfoList {
	ret0 := C.g_file_attribute_info_list_ref(list.native())
	return wrapFileAttributeInfoList(ret0)
}

// Unref is a wrapper around g_file_attribute_info_list_unref().
func (list FileAttributeInfoList) Unref() {
	C.g_file_attribute_info_list_unref(list.native())
}

// Struct FileAttributeInfo
type FileAttributeInfo struct {
	Ptr unsafe.Pointer
}

func (v FileAttributeInfo) native() *C.GFileAttributeInfo {
	return (*C.GFileAttributeInfo)(v.Ptr)
}
func wrapFileAttributeInfo(p *C.GFileAttributeInfo) FileAttributeInfo {
	return FileAttributeInfo{unsafe.Pointer(p)}
}
func WrapFileAttributeInfo(p unsafe.Pointer) FileAttributeInfo {
	return FileAttributeInfo{p}
}

// Object FileInputStream
type FileInputStream struct {
	InputStream
}

func (v FileInputStream) native() *C.GFileInputStream {
	return (*C.GFileInputStream)(v.Ptr)
}
func wrapFileInputStream(p *C.GFileInputStream) (v FileInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileInputStream(p unsafe.Pointer) (v FileInputStream) {
	v.Ptr = p
	return
}
func (v FileInputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// QueryInfo is a wrapper around g_file_input_stream_query_info().
func (stream FileInputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_input_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoFinish is a wrapper around g_file_input_stream_query_info_finish().
func (stream FileInputStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_input_stream_query_info_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

type BusType int

const (
	BusTypeStarter BusType = -1
	BusTypeNone            = 0
	BusTypeSystem          = 1
	BusTypeSession         = 2
)

type ConverterResult int

const (
	ConverterResultError     ConverterResult = 0
	ConverterResultConverted                 = 1
	ConverterResultFinished                  = 2
	ConverterResultFlushed                   = 3
)

type CredentialsType int

const (
	CredentialsTypeInvalid             CredentialsType = 0
	CredentialsTypeLinuxUcred                          = 1
	CredentialsTypeFreebsdCmsgcred                     = 2
	CredentialsTypeOpenbsdSockpeercred                 = 3
	CredentialsTypeSolarisUcred                        = 4
	CredentialsTypeNetbsdUnpcbid                       = 5
)

type DBusError int

const (
	DBusErrorFailed                        DBusError = 0
	DBusErrorNoMemory                                = 1
	DBusErrorServiceUnknown                          = 2
	DBusErrorNameHasNoOwner                          = 3
	DBusErrorNoReply                                 = 4
	DBusErrorIoError                                 = 5
	DBusErrorBadAddress                              = 6
	DBusErrorNotSupported                            = 7
	DBusErrorLimitsExceeded                          = 8
	DBusErrorAccessDenied                            = 9
	DBusErrorAuthFailed                              = 10
	DBusErrorNoServer                                = 11
	DBusErrorTimeout                                 = 12
	DBusErrorNoNetwork                               = 13
	DBusErrorAddressInUse                            = 14
	DBusErrorDisconnected                            = 15
	DBusErrorInvalidArgs                             = 16
	DBusErrorFileNotFound                            = 17
	DBusErrorFileExists                              = 18
	DBusErrorUnknownMethod                           = 19
	DBusErrorTimedOut                                = 20
	DBusErrorMatchRuleNotFound                       = 21
	DBusErrorMatchRuleInvalid                        = 22
	DBusErrorSpawnExecFailed                         = 23
	DBusErrorSpawnForkFailed                         = 24
	DBusErrorSpawnChildExited                        = 25
	DBusErrorSpawnChildSignaled                      = 26
	DBusErrorSpawnFailed                             = 27
	DBusErrorSpawnSetupFailed                        = 28
	DBusErrorSpawnConfigInvalid                      = 29
	DBusErrorSpawnServiceInvalid                     = 30
	DBusErrorSpawnServiceNotFound                    = 31
	DBusErrorSpawnPermissionsInvalid                 = 32
	DBusErrorSpawnFileInvalid                        = 33
	DBusErrorSpawnNoMemory                           = 34
	DBusErrorUnixProcessIdUnknown                    = 35
	DBusErrorInvalidSignature                        = 36
	DBusErrorInvalidFileContent                      = 37
	DBusErrorSelinuxSecurityContextUnknown           = 38
	DBusErrorAdtAuditDataUnknown                     = 39
	DBusErrorObjectPathInUse                         = 40
	DBusErrorUnknownObject                           = 41
	DBusErrorUnknownInterface                        = 42
	DBusErrorUnknownProperty                         = 43
	DBusErrorPropertyReadOnly                        = 44
)

type DBusMessageByteOrder int

const (
	DBusMessageByteOrderBigEndian    DBusMessageByteOrder = 66
	DBusMessageByteOrderLittleEndian                      = 108
)

type DBusMessageHeaderField int

const (
	DBusMessageHeaderFieldInvalid     DBusMessageHeaderField = 0
	DBusMessageHeaderFieldPath                               = 1
	DBusMessageHeaderFieldInterface                          = 2
	DBusMessageHeaderFieldMember                             = 3
	DBusMessageHeaderFieldErrorName                          = 4
	DBusMessageHeaderFieldReplySerial                        = 5
	DBusMessageHeaderFieldDestination                        = 6
	DBusMessageHeaderFieldSender                             = 7
	DBusMessageHeaderFieldSignature                          = 8
	DBusMessageHeaderFieldNumUnixFds                         = 9
)

type DBusMessageType int

const (
	DBusMessageTypeInvalid      DBusMessageType = 0
	DBusMessageTypeMethodCall                   = 1
	DBusMessageTypeMethodReturn                 = 2
	DBusMessageTypeError                        = 3
	DBusMessageTypeSignal                       = 4
)

type DataStreamByteOrder int

const (
	DataStreamByteOrderBigEndian    DataStreamByteOrder = 0
	DataStreamByteOrderLittleEndian                     = 1
	DataStreamByteOrderHostEndian                       = 2
)

type DataStreamNewlineType int

const (
	DataStreamNewlineTypeLf   DataStreamNewlineType = 0
	DataStreamNewlineTypeCr                         = 1
	DataStreamNewlineTypeCrLf                       = 2
	DataStreamNewlineTypeAny                        = 3
)

type DriveStartStopType int

const (
	DriveStartStopTypeUnknown   DriveStartStopType = 0
	DriveStartStopTypeShutdown                     = 1
	DriveStartStopTypeNetwork                      = 2
	DriveStartStopTypeMultidisk                    = 3
	DriveStartStopTypePassword                     = 4
)

type EmblemOrigin int

const (
	EmblemOriginUnknown      EmblemOrigin = 0
	EmblemOriginDevice                    = 1
	EmblemOriginLivemetadata              = 2
	EmblemOriginTag                       = 3
)

type FileAttributeStatus int

const (
	FileAttributeStatusUnset        FileAttributeStatus = 0
	FileAttributeStatusSet                              = 1
	FileAttributeStatusErrorSetting                     = 2
)

type FileAttributeType int

const (
	FileAttributeTypeInvalid    FileAttributeType = 0
	FileAttributeTypeString                       = 1
	FileAttributeTypeByteString                   = 2
	FileAttributeTypeBoolean                      = 3
	FileAttributeTypeUint32                       = 4
	FileAttributeTypeInt32                        = 5
	FileAttributeTypeUint64                       = 6
	FileAttributeTypeInt64                        = 7
	FileAttributeTypeObject                       = 8
	FileAttributeTypeStringv                      = 9
)

type FileMonitorEvent int

const (
	FileMonitorEventChanged          FileMonitorEvent = 0
	FileMonitorEventChangesDoneHint                   = 1
	FileMonitorEventDeleted                           = 2
	FileMonitorEventCreated                           = 3
	FileMonitorEventAttributeChanged                  = 4
	FileMonitorEventPreUnmount                        = 5
	FileMonitorEventUnmounted                         = 6
	FileMonitorEventMoved                             = 7
	FileMonitorEventRenamed                           = 8
	FileMonitorEventMovedIn                           = 9
	FileMonitorEventMovedOut                          = 10
)

type FileType int

const (
	FileTypeUnknown      FileType = 0
	FileTypeRegular               = 1
	FileTypeDirectory             = 2
	FileTypeSymbolicLink          = 3
	FileTypeSpecial               = 4
	FileTypeShortcut              = 5
	FileTypeMountable             = 6
)

type FilesystemPreviewType int

const (
	FilesystemPreviewTypeIfAlways FilesystemPreviewType = 0
	FilesystemPreviewTypeIfLocal                        = 1
	FilesystemPreviewTypeNever                          = 2
)

type IOErrorEnum int

const (
	IOErrorEnumFailed             IOErrorEnum = 0
	IOErrorEnumNotFound                       = 1
	IOErrorEnumExists                         = 2
	IOErrorEnumIsDirectory                    = 3
	IOErrorEnumNotDirectory                   = 4
	IOErrorEnumNotEmpty                       = 5
	IOErrorEnumNotRegularFile                 = 6
	IOErrorEnumNotSymbolicLink                = 7
	IOErrorEnumNotMountableFile               = 8
	IOErrorEnumFilenameTooLong                = 9
	IOErrorEnumInvalidFilename                = 10
	IOErrorEnumTooManyLinks                   = 11
	IOErrorEnumNoSpace                        = 12
	IOErrorEnumInvalidArgument                = 13
	IOErrorEnumPermissionDenied               = 14
	IOErrorEnumNotSupported                   = 15
	IOErrorEnumNotMounted                     = 16
	IOErrorEnumAlreadyMounted                 = 17
	IOErrorEnumClosed                         = 18
	IOErrorEnumCancelled                      = 19
	IOErrorEnumPending                        = 20
	IOErrorEnumReadOnly                       = 21
	IOErrorEnumCantCreateBackup               = 22
	IOErrorEnumWrongEtag                      = 23
	IOErrorEnumTimedOut                       = 24
	IOErrorEnumWouldRecurse                   = 25
	IOErrorEnumBusy                           = 26
	IOErrorEnumWouldBlock                     = 27
	IOErrorEnumHostNotFound                   = 28
	IOErrorEnumWouldMerge                     = 29
	IOErrorEnumFailedHandled                  = 30
	IOErrorEnumTooManyOpenFiles               = 31
	IOErrorEnumNotInitialized                 = 32
	IOErrorEnumAddressInUse                   = 33
	IOErrorEnumPartialInput                   = 34
	IOErrorEnumInvalidData                    = 35
	IOErrorEnumDbusError                      = 36
	IOErrorEnumHostUnreachable                = 37
	IOErrorEnumNetworkUnreachable             = 38
	IOErrorEnumConnectionRefused              = 39
	IOErrorEnumProxyFailed                    = 40
	IOErrorEnumProxyAuthFailed                = 41
	IOErrorEnumProxyNeedAuth                  = 42
	IOErrorEnumProxyNotAllowed                = 43
	IOErrorEnumBrokenPipe                     = 44
	IOErrorEnumConnectionClosed               = 44
	IOErrorEnumNotConnected                   = 45
	IOErrorEnumMessageTooLarge                = 46
)

type IOModuleScopeFlags int

const (
	IOModuleScopeFlagsNone            IOModuleScopeFlags = 0
	IOModuleScopeFlagsBlockDuplicates                    = 1
)

type MountOperationResult int

const (
	MountOperationResultHandled   MountOperationResult = 0
	MountOperationResultAborted                        = 1
	MountOperationResultUnhandled                      = 2
)

type NetworkConnectivity int

const (
	NetworkConnectivityLocal   NetworkConnectivity = 1
	NetworkConnectivityLimited                     = 2
	NetworkConnectivityPortal                      = 3
	NetworkConnectivityFull                        = 4
)

type NotificationPriority int

const (
	NotificationPriorityNormal NotificationPriority = 0
	NotificationPriorityLow                         = 1
	NotificationPriorityHigh                        = 2
	NotificationPriorityUrgent                      = 3
)

type PasswordSave int

const (
	PasswordSaveNever       PasswordSave = 0
	PasswordSaveForSession               = 1
	PasswordSavePermanently              = 2
)

type ResolverError int

const (
	ResolverErrorNotFound         ResolverError = 0
	ResolverErrorTemporaryFailure               = 1
	ResolverErrorInternal                       = 2
)

type ResolverRecordType int

const (
	ResolverRecordTypeSrv ResolverRecordType = 1
	ResolverRecordTypeMx                     = 2
	ResolverRecordTypeTxt                    = 3
	ResolverRecordTypeSoa                    = 4
	ResolverRecordTypeNs                     = 5
)

type ResourceError int

const (
	ResourceErrorNotFound ResourceError = 0
	ResourceErrorInternal               = 1
)

type SocketClientEvent int

const (
	SocketClientEventResolving        SocketClientEvent = 0
	SocketClientEventResolved                           = 1
	SocketClientEventConnecting                         = 2
	SocketClientEventConnected                          = 3
	SocketClientEventProxyNegotiating                   = 4
	SocketClientEventProxyNegotiated                    = 5
	SocketClientEventTlsHandshaking                     = 6
	SocketClientEventTlsHandshaked                      = 7
	SocketClientEventComplete                           = 8
)

type SocketFamily int

const (
	SocketFamilyInvalid SocketFamily = 0
	SocketFamilyUnix                 = 1
	SocketFamilyIpv4                 = 2
	SocketFamilyIpv6                 = 10
)

type SocketListenerEvent int

const (
	SocketListenerEventBinding   SocketListenerEvent = 0
	SocketListenerEventBound                         = 1
	SocketListenerEventListening                     = 2
	SocketListenerEventListened                      = 3
)

type SocketProtocol int

const (
	SocketProtocolUnknown SocketProtocol = -1
	SocketProtocolDefault                = 0
	SocketProtocolTcp                    = 6
	SocketProtocolUdp                    = 17
	SocketProtocolSctp                   = 132
)

type SocketType int

const (
	SocketTypeInvalid   SocketType = 0
	SocketTypeStream               = 1
	SocketTypeDatagram             = 2
	SocketTypeSeqpacket            = 3
)

type TlsAuthenticationMode int

const (
	TlsAuthenticationModeNone      TlsAuthenticationMode = 0
	TlsAuthenticationModeRequested                       = 1
	TlsAuthenticationModeRequired                        = 2
)

type TlsCertificateRequestFlags int

const (
	TlsCertificateRequestFlagsNone TlsCertificateRequestFlags = 0
)

type TlsDatabaseLookupFlags int

const (
	TlsDatabaseLookupFlagsNone    TlsDatabaseLookupFlags = 0
	TlsDatabaseLookupFlagsKeypair                        = 1
)

type TlsError int

const (
	TlsErrorUnavailable         TlsError = 0
	TlsErrorMisc                         = 1
	TlsErrorBadCertificate               = 2
	TlsErrorNotTls                       = 3
	TlsErrorHandshake                    = 4
	TlsErrorCertificateRequired          = 5
	TlsErrorEof                          = 6
)

type TlsInteractionResult int

const (
	TlsInteractionResultUnhandled TlsInteractionResult = 0
	TlsInteractionResultHandled                        = 1
	TlsInteractionResultFailed                         = 2
)

type TlsRehandshakeMode int

const (
	TlsRehandshakeModeNever    TlsRehandshakeMode = 0
	TlsRehandshakeModeSafely                      = 1
	TlsRehandshakeModeUnsafely                    = 2
)

type UnixSocketAddressType int

const (
	UnixSocketAddressTypeInvalid        UnixSocketAddressType = 0
	UnixSocketAddressTypeAnonymous                            = 1
	UnixSocketAddressTypePath                                 = 2
	UnixSocketAddressTypeAbstract                             = 3
	UnixSocketAddressTypeAbstractPadded                       = 4
)

type ZlibCompressorFormat int

const (
	ZlibCompressorFormatZlib ZlibCompressorFormat = 0
	ZlibCompressorFormatGzip                      = 1
	ZlibCompressorFormatRaw                       = 2
)

type AppInfoCreateFlags int

const (
	AppInfoCreateFlagsNone                        AppInfoCreateFlags = 0
	AppInfoCreateFlagsNeedsTerminal                                  = 1
	AppInfoCreateFlagsSupportsUris                                   = 2
	AppInfoCreateFlagsSupportsStartupNotification                    = 4
)

type ApplicationFlags int

const (
	ApplicationFlagsFlagsNone          ApplicationFlags = 0
	ApplicationFlagsIsService                           = 1
	ApplicationFlagsIsLauncher                          = 2
	ApplicationFlagsHandlesOpen                         = 4
	ApplicationFlagsHandlesCommandLine                  = 8
	ApplicationFlagsSendEnvironment                     = 16
	ApplicationFlagsNonUnique                           = 32
	ApplicationFlagsCanOverrideAppId                    = 64
)

type AskPasswordFlags int

const (
	AskPasswordFlagsNeedPassword       AskPasswordFlags = 1
	AskPasswordFlagsNeedUsername                        = 2
	AskPasswordFlagsNeedDomain                          = 4
	AskPasswordFlagsSavingSupported                     = 8
	AskPasswordFlagsAnonymousSupported                  = 16
)

type BusNameOwnerFlags int

const (
	BusNameOwnerFlagsNone             BusNameOwnerFlags = 0
	BusNameOwnerFlagsAllowReplacement                   = 1
	BusNameOwnerFlagsReplace                            = 2
)

type BusNameWatcherFlags int

const (
	BusNameWatcherFlagsNone      BusNameWatcherFlags = 0
	BusNameWatcherFlagsAutoStart                     = 1
)

type ConverterFlags int

const (
	ConverterFlagsNone       ConverterFlags = 0
	ConverterFlagsInputAtEnd                = 1
	ConverterFlagsFlush                     = 2
)

type DBusCallFlags int

const (
	DBusCallFlagsNone                          DBusCallFlags = 0
	DBusCallFlagsNoAutoStart                                 = 1
	DBusCallFlagsAllowInteractiveAuthorization               = 2
)

type DBusCapabilityFlags int

const (
	DBusCapabilityFlagsNone          DBusCapabilityFlags = 0
	DBusCapabilityFlagsUnixFdPassing                     = 1
)

type DBusConnectionFlags int

const (
	DBusConnectionFlagsNone                         DBusConnectionFlags = 0
	DBusConnectionFlagsAuthenticationClient                             = 1
	DBusConnectionFlagsAuthenticationServer                             = 2
	DBusConnectionFlagsAuthenticationAllowAnonymous                     = 4
	DBusConnectionFlagsMessageBusConnection                             = 8
	DBusConnectionFlagsDelayMessageProcessing                           = 16
)

type DBusInterfaceSkeletonFlags int

const (
	DBusInterfaceSkeletonFlagsNone                            DBusInterfaceSkeletonFlags = 0
	DBusInterfaceSkeletonFlagsHandleMethodInvocationsInThread                            = 1
)

type DBusMessageFlags int

const (
	DBusMessageFlagsNone                          DBusMessageFlags = 0
	DBusMessageFlagsNoReplyExpected                                = 1
	DBusMessageFlagsNoAutoStart                                    = 2
	DBusMessageFlagsAllowInteractiveAuthorization                  = 4
)

type DBusObjectManagerClientFlags int

const (
	DBusObjectManagerClientFlagsNone           DBusObjectManagerClientFlags = 0
	DBusObjectManagerClientFlagsDoNotAutoStart                              = 1
)

type DBusPropertyInfoFlags int

const (
	DBusPropertyInfoFlagsNone     DBusPropertyInfoFlags = 0
	DBusPropertyInfoFlagsReadable                       = 1
	DBusPropertyInfoFlagsWritable                       = 2
)

type DBusProxyFlags int

const (
	DBusProxyFlagsNone                         DBusProxyFlags = 0
	DBusProxyFlagsDoNotLoadProperties                         = 1
	DBusProxyFlagsDoNotConnectSignals                         = 2
	DBusProxyFlagsDoNotAutoStart                              = 4
	DBusProxyFlagsGetInvalidatedProperties                    = 8
	DBusProxyFlagsDoNotAutoStartAtConstruction                = 16
)

type DBusSendMessageFlags int

const (
	DBusSendMessageFlagsNone           DBusSendMessageFlags = 0
	DBusSendMessageFlagsPreserveSerial                      = 1
)

type DBusServerFlags int

const (
	DBusServerFlagsNone                         DBusServerFlags = 0
	DBusServerFlagsRunInThread                                  = 1
	DBusServerFlagsAuthenticationAllowAnonymous                 = 2
)

type DBusSignalFlags int

const (
	DBusSignalFlagsNone               DBusSignalFlags = 0
	DBusSignalFlagsNoMatchRule                        = 1
	DBusSignalFlagsMatchArg0Namespace                 = 2
	DBusSignalFlagsMatchArg0Path                      = 4
)

type DBusSubtreeFlags int

const (
	DBusSubtreeFlagsNone                        DBusSubtreeFlags = 0
	DBusSubtreeFlagsDispatchToUnenumeratedNodes                  = 1
)

type DriveStartFlags int

const (
	DriveStartFlagsNone DriveStartFlags = 0
)

type FileAttributeInfoFlags int

const (
	FileAttributeInfoFlagsNone          FileAttributeInfoFlags = 0
	FileAttributeInfoFlagsCopyWithFile                         = 1
	FileAttributeInfoFlagsCopyWhenMoved                        = 2
)

type FileCopyFlags int

const (
	FileCopyFlagsNone               FileCopyFlags = 0
	FileCopyFlagsOverwrite                        = 1
	FileCopyFlagsBackup                           = 2
	FileCopyFlagsNofollowSymlinks                 = 4
	FileCopyFlagsAllMetadata                      = 8
	FileCopyFlagsNoFallbackForMove                = 16
	FileCopyFlagsTargetDefaultPerms               = 32
)

type FileCreateFlags int

const (
	FileCreateFlagsNone               FileCreateFlags = 0
	FileCreateFlagsPrivate                            = 1
	FileCreateFlagsReplaceDestination                 = 2
)

type FileMeasureFlags int

const (
	FileMeasureFlagsNone           FileMeasureFlags = 0
	FileMeasureFlagsReportAnyError                  = 2
	FileMeasureFlagsApparentSize                    = 4
	FileMeasureFlagsNoXdev                          = 8
)

type FileMonitorFlags int

const (
	FileMonitorFlagsNone           FileMonitorFlags = 0
	FileMonitorFlagsWatchMounts                     = 1
	FileMonitorFlagsSendMoved                       = 2
	FileMonitorFlagsWatchHardLinks                  = 4
	FileMonitorFlagsWatchMoves                      = 8
)

type FileQueryInfoFlags int

const (
	FileQueryInfoFlagsNone             FileQueryInfoFlags = 0
	FileQueryInfoFlagsNofollowSymlinks                    = 1
)

type IOStreamSpliceFlags int

const (
	IOStreamSpliceFlagsNone         IOStreamSpliceFlags = 0
	IOStreamSpliceFlagsCloseStream1                     = 1
	IOStreamSpliceFlagsCloseStream2                     = 2
	IOStreamSpliceFlagsWaitForBoth                      = 4
)

type MountMountFlags int

const (
	MountMountFlagsNone MountMountFlags = 0
)

type MountUnmountFlags int

const (
	MountUnmountFlagsNone  MountUnmountFlags = 0
	MountUnmountFlagsForce                   = 1
)

type OutputStreamSpliceFlags int

const (
	OutputStreamSpliceFlagsNone        OutputStreamSpliceFlags = 0
	OutputStreamSpliceFlagsCloseSource                         = 1
	OutputStreamSpliceFlagsCloseTarget                         = 2
)

type ResourceFlags int

const (
	ResourceFlagsNone       ResourceFlags = 0
	ResourceFlagsCompressed               = 1
)

type ResourceLookupFlags int

const (
	ResourceLookupFlagsNone ResourceLookupFlags = 0
)

type SettingsBindFlags int

const (
	SettingsBindFlagsDefault       SettingsBindFlags = 0
	SettingsBindFlagsGet                             = 1
	SettingsBindFlagsSet                             = 2
	SettingsBindFlagsNoSensitivity                   = 4
	SettingsBindFlagsGetNoChanges                    = 8
	SettingsBindFlagsInvertBoolean                   = 16
)

type SocketMsgFlags int

const (
	SocketMsgFlagsNone      SocketMsgFlags = 0
	SocketMsgFlagsOob                      = 1
	SocketMsgFlagsPeek                     = 2
	SocketMsgFlagsDontroute                = 4
)

type SubprocessFlags int

const (
	SubprocessFlagsNone          SubprocessFlags = 0
	SubprocessFlagsStdinPipe                     = 1
	SubprocessFlagsStdinInherit                  = 2
	SubprocessFlagsStdoutPipe                    = 4
	SubprocessFlagsStdoutSilence                 = 8
	SubprocessFlagsStderrPipe                    = 16
	SubprocessFlagsStderrSilence                 = 32
	SubprocessFlagsStderrMerge                   = 64
	SubprocessFlagsInheritFds                    = 128
)

type TestDBusFlags int

const (
	TestDBusFlagsNone TestDBusFlags = 0
)

type TlsCertificateFlags int

const (
	TlsCertificateFlagsUnknownCa    TlsCertificateFlags = 1
	TlsCertificateFlagsBadIdentity                      = 2
	TlsCertificateFlagsNotActivated                     = 4
	TlsCertificateFlagsExpired                          = 8
	TlsCertificateFlagsRevoked                          = 16
	TlsCertificateFlagsInsecure                         = 32
	TlsCertificateFlagsGenericError                     = 64
	TlsCertificateFlagsValidateAll                      = 127
)

type TlsDatabaseVerifyFlags int

const (
	TlsDatabaseVerifyFlagsNone TlsDatabaseVerifyFlags = 0
)

type TlsPasswordFlags int

const (
	TlsPasswordFlagsNone      TlsPasswordFlags = 0
	TlsPasswordFlagsRetry                      = 2
	TlsPasswordFlagsManyTries                  = 4
	TlsPasswordFlagsFinalTry                   = 8
)
