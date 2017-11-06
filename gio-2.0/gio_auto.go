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

// GetPath is a wrapper around g_file_get_path().
func (file File) GetPath() string {
	ret0 := C.g_file_get_path(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
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
