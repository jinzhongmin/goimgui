package gi

//#include <stdio.h>
//#include <stdlib.h>
//const void* __ptr_zero = NULL;
//const int __i32_zero = 0;
//const int __i32_one = 1;
//const int __i32_minus_one = -1;
//const int __i32_false = 0;
//const int __i32_true = 1;
//const unsigned int __u32_zero = 0;
//const float __f32_zero = 0.0f; const float* _f32_zero = &__f32_zero;
//const float __f32_minus_one = -1.0f; const float* _f32_minus_one = &__f32_minus_one;
//const struct __ImVec2_zero__{float x; float y;} __imvec2_zero = {0, 0};
//const struct __ImVec2_ones__{float x; float y;} __imvec2_ones = {1, 1};
//const struct __ImVec2_10__{float x; float y;} __imvec2_10 = {1, 0};
//const struct __ImVec2_01__{float x; float y;} __imvec2_01 = {0, 1};
//const struct __ImVec4_ones__{float x; float y;float z; float w;} __imvec4_ones ={1, 1, 1, 1};
//const struct __ImVec4_zero__{float x; float y;float z; float w;} __imvec4_zero ={0, 0, 0, 0};
import "C"
import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/usf"
)

// flags define
type (
	ImGuiPlotType                uint32
	ImGuiTableFlags              uint32
	ImGuiTreeNodeFlags           uint32
	ImGuiWindowDockStyleCol      uint32
	ImGuiInputSource             uint32
	ImGuiDockNodeFlags           uint32
	ImGuiDragDropFlags           uint32
	ImGuiKey                     uint32
	ImGuiSelectableFlagsPrivate_ uint32
	ImGuiTableBgTarget           uint32
	ImGuiColorEditFlags          uint32
	ImGuiFocusedFlags            uint32
	ImGuiCol                     uint32
	ImGuiNavInput                uint32
	ImGuiNextWindowDataFlags     uint32
	ImGuiLocKey                  uint32
	ImGuiLayoutType              uint32
	ImGuiMouseButton             uint32
	ImGuiNavHighlightFlags       uint32
	ImGuiNextItemDataFlags       uint32
	ImGuiPopupPositionPolicy     uint32
	ImGuiSortDirection           uint32
	ImGuiHoveredFlagsPrivate_    uint32
	ImGuiButtonFlags             uint32
	ImGuiItemStatusFlags         uint32
	ImGuiButtonFlagsPrivate_     uint32
	ImGuiLogType                 uint32
	ImGuiPopupFlags              uint32
	ImGuiInputTextFlags          uint32
	ImGuiBackendFlags            uint32
	ImGuiDir                     int32
	ImGuiTreeNodeFlagsPrivate_   uint32
	ImGuiAxis                    int32
	ImGuiActivateFlags           uint32
	ImGuiFocusRequestFlags       uint32
	ImGuiTabBarFlags             uint32
	ImGuiTableColumnFlags        uint32
	ImDrawFlags                  uint32
	ImGuiDockNodeState           uint32
	ImGuiItemFlags               uint32
	ImGuiOldColumnFlags          uint32
	ImGuiComboFlagsPrivate_      uint32
	ImGuiCond                    uint32
	ImGuiDebugLogFlags           uint32
	ImGuiTableRowFlags           uint32
	ImGuiWindowFlags             uint32
	ImDrawListFlags              uint32
	ImGuiConfigFlags             uint32
	ImGuiMouseSource             uint32
	ImGuiSelectableFlags         uint32
	ImGuiSliderFlags             uint32
	ImGuiTabItemFlagsPrivate_    uint32
	ImGuiTabItemFlags            uint32
	ImFontAtlasFlags             uint32
	ImGuiContextHookType         uint32
	ImGuiDataType                uint32
	ImGuiTextFlags               uint32
	ImGuiTooltipFlags            uint32
	ImGuiComboFlags              uint32
	ImGuiDockNodeFlagsPrivate_   int32
	ImGuiSeparatorFlags          uint32
	ImGuiDataTypePrivate_        uint32
	ImGuiInputEventType          uint32
	ImGuiMouseCursor             int32
	ImGuiNavLayer                uint32
	ImGuiSliderFlagsPrivate_     uint32
	ImGuiStyleVar                uint32
	ImGuiViewportFlags           uint32
	ImGuiDataAuthority           uint32
	ImGuiInputFlags              uint32
	ImGuiInputTextFlagsPrivate_  uint32
	ImGuiNavMoveFlags            uint32
	ImGuiScrollFlags             uint32
	ImGuiTabBarFlagsPrivate_     uint32
	ImGuiHoveredFlags            uint32
)

const (
	ImGuiPlotType_Histogram                                      ImGuiPlotType                = 0x1
	ImGuiPlotType_Lines                                          ImGuiPlotType                = 0x0
	ImGuiTableFlags_Borders                                      ImGuiTableFlags              = 0x780
	ImGuiTableFlags_BordersH                                     ImGuiTableFlags              = 0x180
	ImGuiTableFlags_BordersInner                                 ImGuiTableFlags              = 0x280
	ImGuiTableFlags_BordersInnerH                                ImGuiTableFlags              = 0x80
	ImGuiTableFlags_BordersInnerV                                ImGuiTableFlags              = 0x200
	ImGuiTableFlags_BordersOuter                                 ImGuiTableFlags              = 0x500
	ImGuiTableFlags_BordersOuterH                                ImGuiTableFlags              = 0x100
	ImGuiTableFlags_BordersOuterV                                ImGuiTableFlags              = 0x400
	ImGuiTableFlags_BordersV                                     ImGuiTableFlags              = 0x600
	ImGuiTableFlags_ContextMenuInBody                            ImGuiTableFlags              = 0x20
	ImGuiTableFlags_Hideable                                     ImGuiTableFlags              = 0x4
	ImGuiTableFlags_NoBordersInBody                              ImGuiTableFlags              = 0x800
	ImGuiTableFlags_NoBordersInBodyUntilResize                   ImGuiTableFlags              = 0x1000
	ImGuiTableFlags_NoClip                                       ImGuiTableFlags              = 0x100000
	ImGuiTableFlags_NoHostExtendX                                ImGuiTableFlags              = 0x10000
	ImGuiTableFlags_NoHostExtendY                                ImGuiTableFlags              = 0x20000
	ImGuiTableFlags_NoKeepColumnsVisible                         ImGuiTableFlags              = 0x40000
	ImGuiTableFlags_NoPadInnerX                                  ImGuiTableFlags              = 0x800000
	ImGuiTableFlags_NoPadOuterX                                  ImGuiTableFlags              = 0x400000
	ImGuiTableFlags_NoSavedSettings                              ImGuiTableFlags              = 0x10
	ImGuiTableFlags_None                                         ImGuiTableFlags              = 0x0
	ImGuiTableFlags_PadOuterX                                    ImGuiTableFlags              = 0x200000
	ImGuiTableFlags_PreciseWidths                                ImGuiTableFlags              = 0x80000
	ImGuiTableFlags_Reorderable                                  ImGuiTableFlags              = 0x2
	ImGuiTableFlags_Resizable                                    ImGuiTableFlags              = 0x1
	ImGuiTableFlags_RowBg                                        ImGuiTableFlags              = 0x40
	ImGuiTableFlags_ScrollX                                      ImGuiTableFlags              = 0x1000000
	ImGuiTableFlags_ScrollY                                      ImGuiTableFlags              = 0x2000000
	ImGuiTableFlags_SizingFixedFit                               ImGuiTableFlags              = 0x2000
	ImGuiTableFlags_SizingFixedSame                              ImGuiTableFlags              = 0x4000
	ImGuiTableFlags_SizingMask_                                  ImGuiTableFlags              = 0xe000
	ImGuiTableFlags_SizingStretchProp                            ImGuiTableFlags              = 0x6000
	ImGuiTableFlags_SizingStretchSame                            ImGuiTableFlags              = 0x8000
	ImGuiTableFlags_SortMulti                                    ImGuiTableFlags              = 0x4000000
	ImGuiTableFlags_SortTristate                                 ImGuiTableFlags              = 0x8000000
	ImGuiTableFlags_Sortable                                     ImGuiTableFlags              = 0x8
	ImGuiTreeNodeFlags_AllowOverlap                              ImGuiTreeNodeFlags           = 0x4
	ImGuiTreeNodeFlags_Bullet                                    ImGuiTreeNodeFlags           = 0x200
	ImGuiTreeNodeFlags_CollapsingHeader                          ImGuiTreeNodeFlags           = 0x1a
	ImGuiTreeNodeFlags_DefaultOpen                               ImGuiTreeNodeFlags           = 0x20
	ImGuiTreeNodeFlags_FramePadding                              ImGuiTreeNodeFlags           = 0x400
	ImGuiTreeNodeFlags_Framed                                    ImGuiTreeNodeFlags           = 0x2
	ImGuiTreeNodeFlags_Leaf                                      ImGuiTreeNodeFlags           = 0x100
	ImGuiTreeNodeFlags_NavLeftJumpsBackHere                      ImGuiTreeNodeFlags           = 0x2000
	ImGuiTreeNodeFlags_NoAutoOpenOnLog                           ImGuiTreeNodeFlags           = 0x10
	ImGuiTreeNodeFlags_NoTreePushOnOpen                          ImGuiTreeNodeFlags           = 0x8
	ImGuiTreeNodeFlags_None                                      ImGuiTreeNodeFlags           = 0x0
	ImGuiTreeNodeFlags_OpenOnArrow                               ImGuiTreeNodeFlags           = 0x80
	ImGuiTreeNodeFlags_OpenOnDoubleClick                         ImGuiTreeNodeFlags           = 0x40
	ImGuiTreeNodeFlags_Selected                                  ImGuiTreeNodeFlags           = 0x1
	ImGuiTreeNodeFlags_SpanAvailWidth                            ImGuiTreeNodeFlags           = 0x800
	ImGuiTreeNodeFlags_SpanFullWidth                             ImGuiTreeNodeFlags           = 0x1000
	ImGuiWindowDockStyleCol_COUNT                                ImGuiWindowDockStyleCol      = 0x6
	ImGuiWindowDockStyleCol_Tab                                  ImGuiWindowDockStyleCol      = 0x1
	ImGuiWindowDockStyleCol_TabActive                            ImGuiWindowDockStyleCol      = 0x3
	ImGuiWindowDockStyleCol_TabHovered                           ImGuiWindowDockStyleCol      = 0x2
	ImGuiWindowDockStyleCol_TabUnfocused                         ImGuiWindowDockStyleCol      = 0x4
	ImGuiWindowDockStyleCol_TabUnfocusedActive                   ImGuiWindowDockStyleCol      = 0x5
	ImGuiWindowDockStyleCol_Text                                 ImGuiWindowDockStyleCol      = 0x0
	ImGuiInputSource_COUNT                                       ImGuiInputSource             = 0x5
	ImGuiInputSource_Clipboard                                   ImGuiInputSource             = 0x4
	ImGuiInputSource_Gamepad                                     ImGuiInputSource             = 0x3
	ImGuiInputSource_Keyboard                                    ImGuiInputSource             = 0x2
	ImGuiInputSource_Mouse                                       ImGuiInputSource             = 0x1
	ImGuiInputSource_None                                        ImGuiInputSource             = 0x0
	ImGuiDockNodeFlags_AutoHideTabBar                            ImGuiDockNodeFlags           = 0x40
	ImGuiDockNodeFlags_KeepAliveOnly                             ImGuiDockNodeFlags           = 0x1
	ImGuiDockNodeFlags_NoDockingInCentralNode                    ImGuiDockNodeFlags           = 0x4
	ImGuiDockNodeFlags_NoResize                                  ImGuiDockNodeFlags           = 0x20
	ImGuiDockNodeFlags_NoSplit                                   ImGuiDockNodeFlags           = 0x10
	ImGuiDockNodeFlags_None                                      ImGuiDockNodeFlags           = 0x0
	ImGuiDockNodeFlags_PassthruCentralNode                       ImGuiDockNodeFlags           = 0x8
	ImGuiDragDropFlags_AcceptBeforeDelivery                      ImGuiDragDropFlags           = 0x400
	ImGuiDragDropFlags_AcceptNoDrawDefaultRect                   ImGuiDragDropFlags           = 0x800
	ImGuiDragDropFlags_AcceptNoPreviewTooltip                    ImGuiDragDropFlags           = 0x1000
	ImGuiDragDropFlags_AcceptPeekOnly                            ImGuiDragDropFlags           = 0xc00
	ImGuiDragDropFlags_None                                      ImGuiDragDropFlags           = 0x0
	ImGuiDragDropFlags_SourceAllowNullID                         ImGuiDragDropFlags           = 0x8
	ImGuiDragDropFlags_SourceAutoExpirePayload                   ImGuiDragDropFlags           = 0x20
	ImGuiDragDropFlags_SourceExtern                              ImGuiDragDropFlags           = 0x10
	ImGuiDragDropFlags_SourceNoDisableHover                      ImGuiDragDropFlags           = 0x2
	ImGuiDragDropFlags_SourceNoHoldToOpenOthers                  ImGuiDragDropFlags           = 0x4
	ImGuiDragDropFlags_SourceNoPreviewTooltip                    ImGuiDragDropFlags           = 0x1
	ImGuiKey_0                                                   ImGuiKey                     = 0x218
	ImGuiKey_1                                                   ImGuiKey                     = 0x219
	ImGuiKey_2                                                   ImGuiKey                     = 0x21a
	ImGuiKey_3                                                   ImGuiKey                     = 0x21b
	ImGuiKey_4                                                   ImGuiKey                     = 0x21c
	ImGuiKey_5                                                   ImGuiKey                     = 0x21d
	ImGuiKey_6                                                   ImGuiKey                     = 0x21e
	ImGuiKey_7                                                   ImGuiKey                     = 0x21f
	ImGuiKey_8                                                   ImGuiKey                     = 0x220
	ImGuiKey_9                                                   ImGuiKey                     = 0x221
	ImGuiKey_A                                                   ImGuiKey                     = 0x222
	ImGuiKey_Apostrophe                                          ImGuiKey                     = 0x248
	ImGuiKey_B                                                   ImGuiKey                     = 0x223
	ImGuiKey_Backslash                                           ImGuiKey                     = 0x250
	ImGuiKey_Backspace                                           ImGuiKey                     = 0x20b
	ImGuiKey_C                                                   ImGuiKey                     = 0x224
	ImGuiKey_COUNT                                               ImGuiKey                     = 0x28c
	ImGuiKey_CapsLock                                            ImGuiKey                     = 0x253
	ImGuiKey_Comma                                               ImGuiKey                     = 0x249
	ImGuiKey_D                                                   ImGuiKey                     = 0x225
	ImGuiKey_Delete                                              ImGuiKey                     = 0x20a
	ImGuiKey_DownArrow                                           ImGuiKey                     = 0x204
	ImGuiKey_E                                                   ImGuiKey                     = 0x226
	ImGuiKey_End                                                 ImGuiKey                     = 0x208
	ImGuiKey_Enter                                               ImGuiKey                     = 0x20d
	ImGuiKey_Equal                                               ImGuiKey                     = 0x24e
	ImGuiKey_Escape                                              ImGuiKey                     = 0x20e
	ImGuiKey_F                                                   ImGuiKey                     = 0x227
	ImGuiKey_F1                                                  ImGuiKey                     = 0x23c
	ImGuiKey_F10                                                 ImGuiKey                     = 0x245
	ImGuiKey_F11                                                 ImGuiKey                     = 0x246
	ImGuiKey_F12                                                 ImGuiKey                     = 0x247
	ImGuiKey_F2                                                  ImGuiKey                     = 0x23d
	ImGuiKey_F3                                                  ImGuiKey                     = 0x23e
	ImGuiKey_F4                                                  ImGuiKey                     = 0x23f
	ImGuiKey_F5                                                  ImGuiKey                     = 0x240
	ImGuiKey_F6                                                  ImGuiKey                     = 0x241
	ImGuiKey_F7                                                  ImGuiKey                     = 0x242
	ImGuiKey_F8                                                  ImGuiKey                     = 0x243
	ImGuiKey_F9                                                  ImGuiKey                     = 0x244
	ImGuiKey_G                                                   ImGuiKey                     = 0x228
	ImGuiKey_GamepadBack                                         ImGuiKey                     = 0x26a
	ImGuiKey_GamepadDpadDown                                     ImGuiKey                     = 0x272
	ImGuiKey_GamepadDpadLeft                                     ImGuiKey                     = 0x26f
	ImGuiKey_GamepadDpadRight                                    ImGuiKey                     = 0x270
	ImGuiKey_GamepadDpadUp                                       ImGuiKey                     = 0x271
	ImGuiKey_GamepadFaceDown                                     ImGuiKey                     = 0x26e
	ImGuiKey_GamepadFaceLeft                                     ImGuiKey                     = 0x26b
	ImGuiKey_GamepadFaceRight                                    ImGuiKey                     = 0x26c
	ImGuiKey_GamepadFaceUp                                       ImGuiKey                     = 0x26d
	ImGuiKey_GamepadL1                                           ImGuiKey                     = 0x273
	ImGuiKey_GamepadL2                                           ImGuiKey                     = 0x275
	ImGuiKey_GamepadL3                                           ImGuiKey                     = 0x277
	ImGuiKey_GamepadLStickDown                                   ImGuiKey                     = 0x27c
	ImGuiKey_GamepadLStickLeft                                   ImGuiKey                     = 0x279
	ImGuiKey_GamepadLStickRight                                  ImGuiKey                     = 0x27a
	ImGuiKey_GamepadLStickUp                                     ImGuiKey                     = 0x27b
	ImGuiKey_GamepadR1                                           ImGuiKey                     = 0x274
	ImGuiKey_GamepadR2                                           ImGuiKey                     = 0x276
	ImGuiKey_GamepadR3                                           ImGuiKey                     = 0x278
	ImGuiKey_GamepadRStickDown                                   ImGuiKey                     = 0x280
	ImGuiKey_GamepadRStickLeft                                   ImGuiKey                     = 0x27d
	ImGuiKey_GamepadRStickRight                                  ImGuiKey                     = 0x27e
	ImGuiKey_GamepadRStickUp                                     ImGuiKey                     = 0x27f
	ImGuiKey_GamepadStart                                        ImGuiKey                     = 0x269
	ImGuiKey_GraveAccent                                         ImGuiKey                     = 0x252
	ImGuiKey_H                                                   ImGuiKey                     = 0x229
	ImGuiKey_Home                                                ImGuiKey                     = 0x207
	ImGuiKey_I                                                   ImGuiKey                     = 0x22a
	ImGuiKey_Insert                                              ImGuiKey                     = 0x209
	ImGuiKey_J                                                   ImGuiKey                     = 0x22b
	ImGuiKey_K                                                   ImGuiKey                     = 0x22c
	ImGuiKey_Keypad0                                             ImGuiKey                     = 0x258
	ImGuiKey_Keypad1                                             ImGuiKey                     = 0x259
	ImGuiKey_Keypad2                                             ImGuiKey                     = 0x25a
	ImGuiKey_Keypad3                                             ImGuiKey                     = 0x25b
	ImGuiKey_Keypad4                                             ImGuiKey                     = 0x25c
	ImGuiKey_Keypad5                                             ImGuiKey                     = 0x25d
	ImGuiKey_Keypad6                                             ImGuiKey                     = 0x25e
	ImGuiKey_Keypad7                                             ImGuiKey                     = 0x25f
	ImGuiKey_Keypad8                                             ImGuiKey                     = 0x260
	ImGuiKey_Keypad9                                             ImGuiKey                     = 0x261
	ImGuiKey_KeypadAdd                                           ImGuiKey                     = 0x266
	ImGuiKey_KeypadDecimal                                       ImGuiKey                     = 0x262
	ImGuiKey_KeypadDivide                                        ImGuiKey                     = 0x263
	ImGuiKey_KeypadEnter                                         ImGuiKey                     = 0x267
	ImGuiKey_KeypadEqual                                         ImGuiKey                     = 0x268
	ImGuiKey_KeypadMultiply                                      ImGuiKey                     = 0x264
	ImGuiKey_KeypadSubtract                                      ImGuiKey                     = 0x265
	ImGuiKey_KeysData_OFFSET                                     ImGuiKey                     = 0x0
	ImGuiKey_KeysData_SIZE                                       ImGuiKey                     = 0x28c
	ImGuiKey_L                                                   ImGuiKey                     = 0x22d
	ImGuiKey_LeftAlt                                             ImGuiKey                     = 0x211
	ImGuiKey_LeftArrow                                           ImGuiKey                     = 0x201
	ImGuiKey_LeftBracket                                         ImGuiKey                     = 0x24f
	ImGuiKey_LeftCtrl                                            ImGuiKey                     = 0x20f
	ImGuiKey_LeftShift                                           ImGuiKey                     = 0x210
	ImGuiKey_LeftSuper                                           ImGuiKey                     = 0x212
	ImGuiKey_M                                                   ImGuiKey                     = 0x22e
	ImGuiKey_Menu                                                ImGuiKey                     = 0x217
	ImGuiKey_Minus                                               ImGuiKey                     = 0x24a
	ImGuiKey_MouseLeft                                           ImGuiKey                     = 0x281
	ImGuiKey_MouseMiddle                                         ImGuiKey                     = 0x283
	ImGuiKey_MouseRight                                          ImGuiKey                     = 0x282
	ImGuiKey_MouseWheelX                                         ImGuiKey                     = 0x286
	ImGuiKey_MouseWheelY                                         ImGuiKey                     = 0x287
	ImGuiKey_MouseX1                                             ImGuiKey                     = 0x284
	ImGuiKey_MouseX2                                             ImGuiKey                     = 0x285
	ImGuiKey_N                                                   ImGuiKey                     = 0x22f
	ImGuiKey_NamedKey_BEGIN                                      ImGuiKey                     = 0x200
	ImGuiKey_NamedKey_COUNT                                      ImGuiKey                     = 0x8c
	ImGuiKey_NamedKey_END                                        ImGuiKey                     = 0x28c
	ImGuiKey_None                                                ImGuiKey                     = 0x0
	ImGuiKey_NumLock                                             ImGuiKey                     = 0x255
	ImGuiKey_O                                                   ImGuiKey                     = 0x230
	ImGuiKey_P                                                   ImGuiKey                     = 0x231
	ImGuiKey_PageDown                                            ImGuiKey                     = 0x206
	ImGuiKey_PageUp                                              ImGuiKey                     = 0x205
	ImGuiKey_Pause                                               ImGuiKey                     = 0x257
	ImGuiKey_Period                                              ImGuiKey                     = 0x24b
	ImGuiKey_PrintScreen                                         ImGuiKey                     = 0x256
	ImGuiKey_Q                                                   ImGuiKey                     = 0x232
	ImGuiKey_R                                                   ImGuiKey                     = 0x233
	ImGuiKey_ReservedForModAlt                                   ImGuiKey                     = 0x28a
	ImGuiKey_ReservedForModCtrl                                  ImGuiKey                     = 0x288
	ImGuiKey_ReservedForModShift                                 ImGuiKey                     = 0x289
	ImGuiKey_ReservedForModSuper                                 ImGuiKey                     = 0x28b
	ImGuiKey_RightAlt                                            ImGuiKey                     = 0x215
	ImGuiKey_RightArrow                                          ImGuiKey                     = 0x202
	ImGuiKey_RightBracket                                        ImGuiKey                     = 0x251
	ImGuiKey_RightCtrl                                           ImGuiKey                     = 0x213
	ImGuiKey_RightShift                                          ImGuiKey                     = 0x214
	ImGuiKey_RightSuper                                          ImGuiKey                     = 0x216
	ImGuiKey_S                                                   ImGuiKey                     = 0x234
	ImGuiKey_ScrollLock                                          ImGuiKey                     = 0x254
	ImGuiKey_Semicolon                                           ImGuiKey                     = 0x24d
	ImGuiKey_Slash                                               ImGuiKey                     = 0x24c
	ImGuiKey_Space                                               ImGuiKey                     = 0x20c
	ImGuiKey_T                                                   ImGuiKey                     = 0x235
	ImGuiKey_Tab                                                 ImGuiKey                     = 0x200
	ImGuiKey_U                                                   ImGuiKey                     = 0x236
	ImGuiKey_UpArrow                                             ImGuiKey                     = 0x203
	ImGuiKey_V                                                   ImGuiKey                     = 0x237
	ImGuiKey_W                                                   ImGuiKey                     = 0x238
	ImGuiKey_X                                                   ImGuiKey                     = 0x239
	ImGuiKey_Y                                                   ImGuiKey                     = 0x23a
	ImGuiKey_Z                                                   ImGuiKey                     = 0x23b
	ImGuiMod_Alt                                                 ImGuiKey                     = 0x4000
	ImGuiMod_Ctrl                                                ImGuiKey                     = 0x1000
	ImGuiMod_Mask_                                               ImGuiKey                     = 0xf800
	ImGuiMod_None                                                ImGuiKey                     = 0x0
	ImGuiMod_Shift                                               ImGuiKey                     = 0x2000
	ImGuiMod_Shortcut                                            ImGuiKey                     = 0x800
	ImGuiMod_Super                                               ImGuiKey                     = 0x8000
	ImGuiSelectableFlags_NoHoldingActiveID                       ImGuiSelectableFlagsPrivate_ = 0x100000
	ImGuiSelectableFlags_NoPadWithHalfSpacing                    ImGuiSelectableFlagsPrivate_ = 0x4000000
	ImGuiSelectableFlags_NoSetKeyOwner                           ImGuiSelectableFlagsPrivate_ = 0x8000000
	ImGuiSelectableFlags_SelectOnClick                           ImGuiSelectableFlagsPrivate_ = 0x400000
	ImGuiSelectableFlags_SelectOnNav                             ImGuiSelectableFlagsPrivate_ = 0x200000
	ImGuiSelectableFlags_SelectOnRelease                         ImGuiSelectableFlagsPrivate_ = 0x800000
	ImGuiSelectableFlags_SetNavIdOnHover                         ImGuiSelectableFlagsPrivate_ = 0x2000000
	ImGuiSelectableFlags_SpanAvailWidth                          ImGuiSelectableFlagsPrivate_ = 0x1000000
	ImGuiTableBgTarget_CellBg                                    ImGuiTableBgTarget           = 0x3
	ImGuiTableBgTarget_None                                      ImGuiTableBgTarget           = 0x0
	ImGuiTableBgTarget_RowBg0                                    ImGuiTableBgTarget           = 0x1
	ImGuiTableBgTarget_RowBg1                                    ImGuiTableBgTarget           = 0x2
	ImGuiColorEditFlags_AlphaBar                                 ImGuiColorEditFlags          = 0x10000
	ImGuiColorEditFlags_AlphaPreview                             ImGuiColorEditFlags          = 0x20000
	ImGuiColorEditFlags_AlphaPreviewHalf                         ImGuiColorEditFlags          = 0x40000
	ImGuiColorEditFlags_DataTypeMask_                            ImGuiColorEditFlags          = 0x1800000
	ImGuiColorEditFlags_DefaultOptions_                          ImGuiColorEditFlags          = 0xa900000
	ImGuiColorEditFlags_DisplayHSV                               ImGuiColorEditFlags          = 0x200000
	ImGuiColorEditFlags_DisplayHex                               ImGuiColorEditFlags          = 0x400000
	ImGuiColorEditFlags_DisplayMask_                             ImGuiColorEditFlags          = 0x700000
	ImGuiColorEditFlags_DisplayRGB                               ImGuiColorEditFlags          = 0x100000
	ImGuiColorEditFlags_Float                                    ImGuiColorEditFlags          = 0x1000000
	ImGuiColorEditFlags_HDR                                      ImGuiColorEditFlags          = 0x80000
	ImGuiColorEditFlags_InputHSV                                 ImGuiColorEditFlags          = 0x10000000
	ImGuiColorEditFlags_InputMask_                               ImGuiColorEditFlags          = 0x18000000
	ImGuiColorEditFlags_InputRGB                                 ImGuiColorEditFlags          = 0x8000000
	ImGuiColorEditFlags_NoAlpha                                  ImGuiColorEditFlags          = 0x2
	ImGuiColorEditFlags_NoBorder                                 ImGuiColorEditFlags          = 0x400
	ImGuiColorEditFlags_NoDragDrop                               ImGuiColorEditFlags          = 0x200
	ImGuiColorEditFlags_NoInputs                                 ImGuiColorEditFlags          = 0x20
	ImGuiColorEditFlags_NoLabel                                  ImGuiColorEditFlags          = 0x80
	ImGuiColorEditFlags_NoOptions                                ImGuiColorEditFlags          = 0x8
	ImGuiColorEditFlags_NoPicker                                 ImGuiColorEditFlags          = 0x4
	ImGuiColorEditFlags_NoSidePreview                            ImGuiColorEditFlags          = 0x100
	ImGuiColorEditFlags_NoSmallPreview                           ImGuiColorEditFlags          = 0x10
	ImGuiColorEditFlags_NoTooltip                                ImGuiColorEditFlags          = 0x40
	ImGuiColorEditFlags_None                                     ImGuiColorEditFlags          = 0x0
	ImGuiColorEditFlags_PickerHueBar                             ImGuiColorEditFlags          = 0x2000000
	ImGuiColorEditFlags_PickerHueWheel                           ImGuiColorEditFlags          = 0x4000000
	ImGuiColorEditFlags_PickerMask_                              ImGuiColorEditFlags          = 0x6000000
	ImGuiColorEditFlags_Uint8                                    ImGuiColorEditFlags          = 0x800000
	ImGuiFocusedFlags_AnyWindow                                  ImGuiFocusedFlags            = 0x4
	ImGuiFocusedFlags_ChildWindows                               ImGuiFocusedFlags            = 0x1
	ImGuiFocusedFlags_DockHierarchy                              ImGuiFocusedFlags            = 0x10
	ImGuiFocusedFlags_NoPopupHierarchy                           ImGuiFocusedFlags            = 0x8
	ImGuiFocusedFlags_None                                       ImGuiFocusedFlags            = 0x0
	ImGuiFocusedFlags_RootAndChildWindows                        ImGuiFocusedFlags            = 0x3
	ImGuiFocusedFlags_RootWindow                                 ImGuiFocusedFlags            = 0x2
	ImGuiCol_Border                                              ImGuiCol                     = 0x5
	ImGuiCol_BorderShadow                                        ImGuiCol                     = 0x6
	ImGuiCol_Button                                              ImGuiCol                     = 0x15
	ImGuiCol_ButtonActive                                        ImGuiCol                     = 0x17
	ImGuiCol_ButtonHovered                                       ImGuiCol                     = 0x16
	ImGuiCol_COUNT                                               ImGuiCol                     = 0x37
	ImGuiCol_CheckMark                                           ImGuiCol                     = 0x12
	ImGuiCol_ChildBg                                             ImGuiCol                     = 0x3
	ImGuiCol_DockingEmptyBg                                      ImGuiCol                     = 0x27
	ImGuiCol_DockingPreview                                      ImGuiCol                     = 0x26
	ImGuiCol_DragDropTarget                                      ImGuiCol                     = 0x32
	ImGuiCol_FrameBg                                             ImGuiCol                     = 0x7
	ImGuiCol_FrameBgActive                                       ImGuiCol                     = 0x9
	ImGuiCol_FrameBgHovered                                      ImGuiCol                     = 0x8
	ImGuiCol_Header                                              ImGuiCol                     = 0x18
	ImGuiCol_HeaderActive                                        ImGuiCol                     = 0x1a
	ImGuiCol_HeaderHovered                                       ImGuiCol                     = 0x19
	ImGuiCol_MenuBarBg                                           ImGuiCol                     = 0xd
	ImGuiCol_ModalWindowDimBg                                    ImGuiCol                     = 0x36
	ImGuiCol_NavHighlight                                        ImGuiCol                     = 0x33
	ImGuiCol_NavWindowingDimBg                                   ImGuiCol                     = 0x35
	ImGuiCol_NavWindowingHighlight                               ImGuiCol                     = 0x34
	ImGuiCol_PlotHistogram                                       ImGuiCol                     = 0x2a
	ImGuiCol_PlotHistogramHovered                                ImGuiCol                     = 0x2b
	ImGuiCol_PlotLines                                           ImGuiCol                     = 0x28
	ImGuiCol_PlotLinesHovered                                    ImGuiCol                     = 0x29
	ImGuiCol_PopupBg                                             ImGuiCol                     = 0x4
	ImGuiCol_ResizeGrip                                          ImGuiCol                     = 0x1e
	ImGuiCol_ResizeGripActive                                    ImGuiCol                     = 0x20
	ImGuiCol_ResizeGripHovered                                   ImGuiCol                     = 0x1f
	ImGuiCol_ScrollbarBg                                         ImGuiCol                     = 0xe
	ImGuiCol_ScrollbarGrab                                       ImGuiCol                     = 0xf
	ImGuiCol_ScrollbarGrabActive                                 ImGuiCol                     = 0x11
	ImGuiCol_ScrollbarGrabHovered                                ImGuiCol                     = 0x10
	ImGuiCol_Separator                                           ImGuiCol                     = 0x1b
	ImGuiCol_SeparatorActive                                     ImGuiCol                     = 0x1d
	ImGuiCol_SeparatorHovered                                    ImGuiCol                     = 0x1c
	ImGuiCol_SliderGrab                                          ImGuiCol                     = 0x13
	ImGuiCol_SliderGrabActive                                    ImGuiCol                     = 0x14
	ImGuiCol_Tab                                                 ImGuiCol                     = 0x21
	ImGuiCol_TabActive                                           ImGuiCol                     = 0x23
	ImGuiCol_TabHovered                                          ImGuiCol                     = 0x22
	ImGuiCol_TabUnfocused                                        ImGuiCol                     = 0x24
	ImGuiCol_TabUnfocusedActive                                  ImGuiCol                     = 0x25
	ImGuiCol_TableBorderLight                                    ImGuiCol                     = 0x2e
	ImGuiCol_TableBorderStrong                                   ImGuiCol                     = 0x2d
	ImGuiCol_TableHeaderBg                                       ImGuiCol                     = 0x2c
	ImGuiCol_TableRowBg                                          ImGuiCol                     = 0x2f
	ImGuiCol_TableRowBgAlt                                       ImGuiCol                     = 0x30
	ImGuiCol_Text                                                ImGuiCol                     = 0x0
	ImGuiCol_TextDisabled                                        ImGuiCol                     = 0x1
	ImGuiCol_TextSelectedBg                                      ImGuiCol                     = 0x31
	ImGuiCol_TitleBg                                             ImGuiCol                     = 0xa
	ImGuiCol_TitleBgActive                                       ImGuiCol                     = 0xb
	ImGuiCol_TitleBgCollapsed                                    ImGuiCol                     = 0xc
	ImGuiCol_WindowBg                                            ImGuiCol                     = 0x2
	ImGuiNavInput_Activate                                       ImGuiNavInput                = 0x0
	ImGuiNavInput_COUNT                                          ImGuiNavInput                = 0x10
	ImGuiNavInput_Cancel                                         ImGuiNavInput                = 0x1
	ImGuiNavInput_DpadDown                                       ImGuiNavInput                = 0x7
	ImGuiNavInput_DpadLeft                                       ImGuiNavInput                = 0x4
	ImGuiNavInput_DpadRight                                      ImGuiNavInput                = 0x5
	ImGuiNavInput_DpadUp                                         ImGuiNavInput                = 0x6
	ImGuiNavInput_FocusNext                                      ImGuiNavInput                = 0xd
	ImGuiNavInput_FocusPrev                                      ImGuiNavInput                = 0xc
	ImGuiNavInput_Input                                          ImGuiNavInput                = 0x2
	ImGuiNavInput_LStickDown                                     ImGuiNavInput                = 0xb
	ImGuiNavInput_LStickLeft                                     ImGuiNavInput                = 0x8
	ImGuiNavInput_LStickRight                                    ImGuiNavInput                = 0x9
	ImGuiNavInput_LStickUp                                       ImGuiNavInput                = 0xa
	ImGuiNavInput_Menu                                           ImGuiNavInput                = 0x3
	ImGuiNavInput_TweakFast                                      ImGuiNavInput                = 0xf
	ImGuiNavInput_TweakSlow                                      ImGuiNavInput                = 0xe
	ImGuiNextWindowDataFlags_HasBgAlpha                          ImGuiNextWindowDataFlags     = 0x40
	ImGuiNextWindowDataFlags_HasCollapsed                        ImGuiNextWindowDataFlags     = 0x8
	ImGuiNextWindowDataFlags_HasContentSize                      ImGuiNextWindowDataFlags     = 0x4
	ImGuiNextWindowDataFlags_HasDock                             ImGuiNextWindowDataFlags     = 0x200
	ImGuiNextWindowDataFlags_HasFocus                            ImGuiNextWindowDataFlags     = 0x20
	ImGuiNextWindowDataFlags_HasPos                              ImGuiNextWindowDataFlags     = 0x1
	ImGuiNextWindowDataFlags_HasScroll                           ImGuiNextWindowDataFlags     = 0x80
	ImGuiNextWindowDataFlags_HasSize                             ImGuiNextWindowDataFlags     = 0x2
	ImGuiNextWindowDataFlags_HasSizeConstraint                   ImGuiNextWindowDataFlags     = 0x10
	ImGuiNextWindowDataFlags_HasViewport                         ImGuiNextWindowDataFlags     = 0x100
	ImGuiNextWindowDataFlags_HasWindowClass                      ImGuiNextWindowDataFlags     = 0x400
	ImGuiNextWindowDataFlags_None                                ImGuiNextWindowDataFlags     = 0x0
	ImGuiLocKey_COUNT                                            ImGuiLocKey                  = 0x9
	ImGuiLocKey_DockingHideTabBar                                ImGuiLocKey                  = 0x8
	ImGuiLocKey_TableResetOrder                                  ImGuiLocKey                  = 0x4
	ImGuiLocKey_TableSizeAllDefault                              ImGuiLocKey                  = 0x3
	ImGuiLocKey_TableSizeAllFit                                  ImGuiLocKey                  = 0x2
	ImGuiLocKey_TableSizeOne                                     ImGuiLocKey                  = 0x1
	ImGuiLocKey_VersionStr                                       ImGuiLocKey                  = 0x0
	ImGuiLocKey_WindowingMainMenuBar                             ImGuiLocKey                  = 0x5
	ImGuiLocKey_WindowingPopup                                   ImGuiLocKey                  = 0x6
	ImGuiLocKey_WindowingUntitled                                ImGuiLocKey                  = 0x7
	ImGuiLayoutType_Horizontal                                   ImGuiLayoutType              = 0x0
	ImGuiLayoutType_Vertical                                     ImGuiLayoutType              = 0x1
	ImGuiMouseButton_COUNT                                       ImGuiMouseButton             = 0x5
	ImGuiMouseButton_Left                                        ImGuiMouseButton             = 0x0
	ImGuiMouseButton_Middle                                      ImGuiMouseButton             = 0x2
	ImGuiMouseButton_Right                                       ImGuiMouseButton             = 0x1
	ImGuiNavHighlightFlags_AlwaysDraw                            ImGuiNavHighlightFlags       = 0x4
	ImGuiNavHighlightFlags_NoRounding                            ImGuiNavHighlightFlags       = 0x8
	ImGuiNavHighlightFlags_None                                  ImGuiNavHighlightFlags       = 0x0
	ImGuiNavHighlightFlags_TypeDefault                           ImGuiNavHighlightFlags       = 0x1
	ImGuiNavHighlightFlags_TypeThin                              ImGuiNavHighlightFlags       = 0x2
	ImGuiNextItemDataFlags_HasOpen                               ImGuiNextItemDataFlags       = 0x2
	ImGuiNextItemDataFlags_HasWidth                              ImGuiNextItemDataFlags       = 0x1
	ImGuiNextItemDataFlags_None                                  ImGuiNextItemDataFlags       = 0x0
	ImGuiPopupPositionPolicy_ComboBox                            ImGuiPopupPositionPolicy     = 0x1
	ImGuiPopupPositionPolicy_Default                             ImGuiPopupPositionPolicy     = 0x0
	ImGuiPopupPositionPolicy_Tooltip                             ImGuiPopupPositionPolicy     = 0x2
	ImGuiSortDirection_Ascending                                 ImGuiSortDirection           = 0x1
	ImGuiSortDirection_Descending                                ImGuiSortDirection           = 0x2
	ImGuiSortDirection_None                                      ImGuiSortDirection           = 0x0
	ImGuiHoveredFlags_AllowedMaskForIsItemHovered                ImGuiHoveredFlagsPrivate_    = 0x1ffa0
	ImGuiHoveredFlags_AllowedMaskForIsWindowHovered              ImGuiHoveredFlagsPrivate_    = 0x18bf
	ImGuiHoveredFlags_DelayMask_                                 ImGuiHoveredFlagsPrivate_    = 0x1e000
	ImGuiButtonFlags_MouseButtonDefault_                         ImGuiButtonFlags             = 0x1
	ImGuiButtonFlags_MouseButtonLeft                             ImGuiButtonFlags             = 0x1
	ImGuiButtonFlags_MouseButtonMask_                            ImGuiButtonFlags             = 0x7
	ImGuiButtonFlags_MouseButtonMiddle                           ImGuiButtonFlags             = 0x4
	ImGuiButtonFlags_MouseButtonRight                            ImGuiButtonFlags             = 0x2
	ImGuiButtonFlags_None                                        ImGuiButtonFlags             = 0x0
	ImGuiItemStatusFlags_Deactivated                             ImGuiItemStatusFlags         = 0x40
	ImGuiItemStatusFlags_Edited                                  ImGuiItemStatusFlags         = 0x4
	ImGuiItemStatusFlags_FocusedByTabbing                        ImGuiItemStatusFlags         = 0x100
	ImGuiItemStatusFlags_HasDeactivated                          ImGuiItemStatusFlags         = 0x20
	ImGuiItemStatusFlags_HasDisplayRect                          ImGuiItemStatusFlags         = 0x2
	ImGuiItemStatusFlags_HoveredRect                             ImGuiItemStatusFlags         = 0x1
	ImGuiItemStatusFlags_HoveredWindow                           ImGuiItemStatusFlags         = 0x80
	ImGuiItemStatusFlags_None                                    ImGuiItemStatusFlags         = 0x0
	ImGuiItemStatusFlags_ToggledOpen                             ImGuiItemStatusFlags         = 0x10
	ImGuiItemStatusFlags_ToggledSelection                        ImGuiItemStatusFlags         = 0x8
	ImGuiItemStatusFlags_Visible                                 ImGuiItemStatusFlags         = 0x200
	ImGuiButtonFlags_AlignTextBaseLine                           ImGuiButtonFlagsPrivate_     = 0x8000
	ImGuiButtonFlags_AllowOverlap                                ImGuiButtonFlagsPrivate_     = 0x1000
	ImGuiButtonFlags_DontClosePopups                             ImGuiButtonFlagsPrivate_     = 0x2000
	ImGuiButtonFlags_FlattenChildren                             ImGuiButtonFlagsPrivate_     = 0x800
	ImGuiButtonFlags_NoHoldingActiveId                           ImGuiButtonFlagsPrivate_     = 0x20000
	ImGuiButtonFlags_NoHoveredOnFocus                            ImGuiButtonFlagsPrivate_     = 0x80000
	ImGuiButtonFlags_NoKeyModifiers                              ImGuiButtonFlagsPrivate_     = 0x10000
	ImGuiButtonFlags_NoNavFocus                                  ImGuiButtonFlagsPrivate_     = 0x40000
	ImGuiButtonFlags_NoSetKeyOwner                               ImGuiButtonFlagsPrivate_     = 0x100000
	ImGuiButtonFlags_NoTestKeyOwner                              ImGuiButtonFlagsPrivate_     = 0x200000
	ImGuiButtonFlags_PressedOnClick                              ImGuiButtonFlagsPrivate_     = 0x10
	ImGuiButtonFlags_PressedOnClickRelease                       ImGuiButtonFlagsPrivate_     = 0x20
	ImGuiButtonFlags_PressedOnClickReleaseAnywhere               ImGuiButtonFlagsPrivate_     = 0x40
	ImGuiButtonFlags_PressedOnDefault_                           ImGuiButtonFlagsPrivate_     = 0x20
	ImGuiButtonFlags_PressedOnDoubleClick                        ImGuiButtonFlagsPrivate_     = 0x100
	ImGuiButtonFlags_PressedOnDragDropHold                       ImGuiButtonFlagsPrivate_     = 0x200
	ImGuiButtonFlags_PressedOnMask_                              ImGuiButtonFlagsPrivate_     = 0x3f0
	ImGuiButtonFlags_PressedOnRelease                            ImGuiButtonFlagsPrivate_     = 0x80
	ImGuiButtonFlags_Repeat                                      ImGuiButtonFlagsPrivate_     = 0x400
	ImGuiLogType_Buffer                                          ImGuiLogType                 = 0x3
	ImGuiLogType_Clipboard                                       ImGuiLogType                 = 0x4
	ImGuiLogType_File                                            ImGuiLogType                 = 0x2
	ImGuiLogType_None                                            ImGuiLogType                 = 0x0
	ImGuiLogType_TTY                                             ImGuiLogType                 = 0x1
	ImGuiPopupFlags_AnyPopup                                     ImGuiPopupFlags              = 0x180
	ImGuiPopupFlags_AnyPopupId                                   ImGuiPopupFlags              = 0x80
	ImGuiPopupFlags_AnyPopupLevel                                ImGuiPopupFlags              = 0x100
	ImGuiPopupFlags_MouseButtonDefault_                          ImGuiPopupFlags              = 0x1
	ImGuiPopupFlags_MouseButtonLeft                              ImGuiPopupFlags              = 0x0
	ImGuiPopupFlags_MouseButtonMask_                             ImGuiPopupFlags              = 0x1f
	ImGuiPopupFlags_MouseButtonMiddle                            ImGuiPopupFlags              = 0x2
	ImGuiPopupFlags_MouseButtonRight                             ImGuiPopupFlags              = 0x1
	ImGuiPopupFlags_NoOpenOverExistingPopup                      ImGuiPopupFlags              = 0x20
	ImGuiPopupFlags_NoOpenOverItems                              ImGuiPopupFlags              = 0x40
	ImGuiPopupFlags_None                                         ImGuiPopupFlags              = 0x0
	ImGuiInputTextFlags_AllowTabInput                            ImGuiInputTextFlags          = 0x400
	ImGuiInputTextFlags_AlwaysOverwrite                          ImGuiInputTextFlags          = 0x2000
	ImGuiInputTextFlags_AutoSelectAll                            ImGuiInputTextFlags          = 0x10
	ImGuiInputTextFlags_CallbackAlways                           ImGuiInputTextFlags          = 0x100
	ImGuiInputTextFlags_CallbackCharFilter                       ImGuiInputTextFlags          = 0x200
	ImGuiInputTextFlags_CallbackCompletion                       ImGuiInputTextFlags          = 0x40
	ImGuiInputTextFlags_CallbackEdit                             ImGuiInputTextFlags          = 0x80000
	ImGuiInputTextFlags_CallbackHistory                          ImGuiInputTextFlags          = 0x80
	ImGuiInputTextFlags_CallbackResize                           ImGuiInputTextFlags          = 0x40000
	ImGuiInputTextFlags_CharsDecimal                             ImGuiInputTextFlags          = 0x1
	ImGuiInputTextFlags_CharsHexadecimal                         ImGuiInputTextFlags          = 0x2
	ImGuiInputTextFlags_CharsNoBlank                             ImGuiInputTextFlags          = 0x8
	ImGuiInputTextFlags_CharsScientific                          ImGuiInputTextFlags          = 0x20000
	ImGuiInputTextFlags_CharsUppercase                           ImGuiInputTextFlags          = 0x4
	ImGuiInputTextFlags_CtrlEnterForNewLine                      ImGuiInputTextFlags          = 0x800
	ImGuiInputTextFlags_EnterReturnsTrue                         ImGuiInputTextFlags          = 0x20
	ImGuiInputTextFlags_EscapeClearsAll                          ImGuiInputTextFlags          = 0x100000
	ImGuiInputTextFlags_NoHorizontalScroll                       ImGuiInputTextFlags          = 0x1000
	ImGuiInputTextFlags_NoUndoRedo                               ImGuiInputTextFlags          = 0x10000
	ImGuiInputTextFlags_None                                     ImGuiInputTextFlags          = 0x0
	ImGuiInputTextFlags_Password                                 ImGuiInputTextFlags          = 0x8000
	ImGuiInputTextFlags_ReadOnly                                 ImGuiInputTextFlags          = 0x4000
	ImGuiBackendFlags_HasGamepad                                 ImGuiBackendFlags            = 0x1
	ImGuiBackendFlags_HasMouseCursors                            ImGuiBackendFlags            = 0x2
	ImGuiBackendFlags_HasMouseHoveredViewport                    ImGuiBackendFlags            = 0x800
	ImGuiBackendFlags_HasSetMousePos                             ImGuiBackendFlags            = 0x4
	ImGuiBackendFlags_None                                       ImGuiBackendFlags            = 0x0
	ImGuiBackendFlags_PlatformHasViewports                       ImGuiBackendFlags            = 0x400
	ImGuiBackendFlags_RendererHasViewports                       ImGuiBackendFlags            = 0x1000
	ImGuiBackendFlags_RendererHasVtxOffset                       ImGuiBackendFlags            = 0x8
	ImGuiDir_COUNT                                               ImGuiDir                     = 0x4
	ImGuiDir_Down                                                ImGuiDir                     = 0x3
	ImGuiDir_Left                                                ImGuiDir                     = 0x0
	ImGuiDir_None                                                ImGuiDir                     = -0x1
	ImGuiDir_Right                                               ImGuiDir                     = 0x1
	ImGuiDir_Up                                                  ImGuiDir                     = 0x2
	ImGuiTreeNodeFlags_ClipLabelForTrailingButton                ImGuiTreeNodeFlagsPrivate_   = 0x100000
	ImGuiTreeNodeFlags_UpsideDownArrow                           ImGuiTreeNodeFlagsPrivate_   = 0x200000
	ImGuiAxis_None                                               ImGuiAxis                    = -0x1
	ImGuiAxis_X                                                  ImGuiAxis                    = 0x0
	ImGuiAxis_Y                                                  ImGuiAxis                    = 0x1
	ImGuiActivateFlags_None                                      ImGuiActivateFlags           = 0x0
	ImGuiActivateFlags_PreferInput                               ImGuiActivateFlags           = 0x1
	ImGuiActivateFlags_PreferTweak                               ImGuiActivateFlags           = 0x2
	ImGuiActivateFlags_TryToPreserveState                        ImGuiActivateFlags           = 0x4
	ImGuiFocusRequestFlags_None                                  ImGuiFocusRequestFlags       = 0x0
	ImGuiFocusRequestFlags_RestoreFocusedChild                   ImGuiFocusRequestFlags       = 0x1
	ImGuiFocusRequestFlags_UnlessBelowModal                      ImGuiFocusRequestFlags       = 0x2
	ImGuiTabBarFlags_AutoSelectNewTabs                           ImGuiTabBarFlags             = 0x2
	ImGuiTabBarFlags_FittingPolicyDefault_                       ImGuiTabBarFlags             = 0x40
	ImGuiTabBarFlags_FittingPolicyMask_                          ImGuiTabBarFlags             = 0xc0
	ImGuiTabBarFlags_FittingPolicyResizeDown                     ImGuiTabBarFlags             = 0x40
	ImGuiTabBarFlags_FittingPolicyScroll                         ImGuiTabBarFlags             = 0x80
	ImGuiTabBarFlags_NoCloseWithMiddleMouseButton                ImGuiTabBarFlags             = 0x8
	ImGuiTabBarFlags_NoTabListScrollingButtons                   ImGuiTabBarFlags             = 0x10
	ImGuiTabBarFlags_NoTooltip                                   ImGuiTabBarFlags             = 0x20
	ImGuiTabBarFlags_None                                        ImGuiTabBarFlags             = 0x0
	ImGuiTabBarFlags_Reorderable                                 ImGuiTabBarFlags             = 0x1
	ImGuiTabBarFlags_TabListPopupButton                          ImGuiTabBarFlags             = 0x4
	ImGuiTableColumnFlags_DefaultHide                            ImGuiTableColumnFlags        = 0x2
	ImGuiTableColumnFlags_DefaultSort                            ImGuiTableColumnFlags        = 0x4
	ImGuiTableColumnFlags_Disabled                               ImGuiTableColumnFlags        = 0x1
	ImGuiTableColumnFlags_IndentDisable                          ImGuiTableColumnFlags        = 0x20000
	ImGuiTableColumnFlags_IndentEnable                           ImGuiTableColumnFlags        = 0x10000
	ImGuiTableColumnFlags_IndentMask_                            ImGuiTableColumnFlags        = 0x30000
	ImGuiTableColumnFlags_IsEnabled                              ImGuiTableColumnFlags        = 0x1000000
	ImGuiTableColumnFlags_IsHovered                              ImGuiTableColumnFlags        = 0x8000000
	ImGuiTableColumnFlags_IsSorted                               ImGuiTableColumnFlags        = 0x4000000
	ImGuiTableColumnFlags_IsVisible                              ImGuiTableColumnFlags        = 0x2000000
	ImGuiTableColumnFlags_NoClip                                 ImGuiTableColumnFlags        = 0x100
	ImGuiTableColumnFlags_NoDirectResize_                        ImGuiTableColumnFlags        = 0x40000000
	ImGuiTableColumnFlags_NoHeaderLabel                          ImGuiTableColumnFlags        = 0x1000
	ImGuiTableColumnFlags_NoHeaderWidth                          ImGuiTableColumnFlags        = 0x2000
	ImGuiTableColumnFlags_NoHide                                 ImGuiTableColumnFlags        = 0x80
	ImGuiTableColumnFlags_NoReorder                              ImGuiTableColumnFlags        = 0x40
	ImGuiTableColumnFlags_NoResize                               ImGuiTableColumnFlags        = 0x20
	ImGuiTableColumnFlags_NoSort                                 ImGuiTableColumnFlags        = 0x200
	ImGuiTableColumnFlags_NoSortAscending                        ImGuiTableColumnFlags        = 0x400
	ImGuiTableColumnFlags_NoSortDescending                       ImGuiTableColumnFlags        = 0x800
	ImGuiTableColumnFlags_None                                   ImGuiTableColumnFlags        = 0x0
	ImGuiTableColumnFlags_PreferSortAscending                    ImGuiTableColumnFlags        = 0x4000
	ImGuiTableColumnFlags_PreferSortDescending                   ImGuiTableColumnFlags        = 0x8000
	ImGuiTableColumnFlags_StatusMask_                            ImGuiTableColumnFlags        = 0xf000000
	ImGuiTableColumnFlags_WidthFixed                             ImGuiTableColumnFlags        = 0x10
	ImGuiTableColumnFlags_WidthMask_                             ImGuiTableColumnFlags        = 0x18
	ImGuiTableColumnFlags_WidthStretch                           ImGuiTableColumnFlags        = 0x8
	ImDrawFlags_Closed                                           ImDrawFlags                  = 0x1
	ImDrawFlags_None                                             ImDrawFlags                  = 0x0
	ImDrawFlags_RoundCornersAll                                  ImDrawFlags                  = 0xf0
	ImDrawFlags_RoundCornersBottom                               ImDrawFlags                  = 0xc0
	ImDrawFlags_RoundCornersBottomLeft                           ImDrawFlags                  = 0x40
	ImDrawFlags_RoundCornersBottomRight                          ImDrawFlags                  = 0x80
	ImDrawFlags_RoundCornersDefault_                             ImDrawFlags                  = 0xf0
	ImDrawFlags_RoundCornersLeft                                 ImDrawFlags                  = 0x50
	ImDrawFlags_RoundCornersMask_                                ImDrawFlags                  = 0x1f0
	ImDrawFlags_RoundCornersNone                                 ImDrawFlags                  = 0x100
	ImDrawFlags_RoundCornersRight                                ImDrawFlags                  = 0xa0
	ImDrawFlags_RoundCornersTop                                  ImDrawFlags                  = 0x30
	ImDrawFlags_RoundCornersTopLeft                              ImDrawFlags                  = 0x10
	ImDrawFlags_RoundCornersTopRight                             ImDrawFlags                  = 0x20
	ImGuiDockNodeState_HostWindowHiddenBecauseSingleWindow       ImGuiDockNodeState           = 0x1
	ImGuiDockNodeState_HostWindowHiddenBecauseWindowsAreResizing ImGuiDockNodeState           = 0x2
	ImGuiDockNodeState_HostWindowVisible                         ImGuiDockNodeState           = 0x3
	ImGuiDockNodeState_Unknown                                   ImGuiDockNodeState           = 0x0
	ImGuiItemFlags_ButtonRepeat                                  ImGuiItemFlags               = 0x2
	ImGuiItemFlags_Disabled                                      ImGuiItemFlags               = 0x4
	ImGuiItemFlags_Inputable                                     ImGuiItemFlags               = 0x400
	ImGuiItemFlags_MixedValue                                    ImGuiItemFlags               = 0x40
	ImGuiItemFlags_NoNav                                         ImGuiItemFlags               = 0x8
	ImGuiItemFlags_NoNavDefaultFocus                             ImGuiItemFlags               = 0x10
	ImGuiItemFlags_NoTabStop                                     ImGuiItemFlags               = 0x1
	ImGuiItemFlags_NoWindowHoverableCheck                        ImGuiItemFlags               = 0x100
	ImGuiItemFlags_None                                          ImGuiItemFlags               = 0x0
	ImGuiItemFlags_ReadOnly                                      ImGuiItemFlags               = 0x80
	ImGuiItemFlags_SelectableDontClosePopup                      ImGuiItemFlags               = 0x20
	ImGuiItemflags_AllowOverlap                                  ImGuiItemFlags               = 0x200
	ImGuiOldColumnFlags_GrowParentContentsSize                   ImGuiOldColumnFlags          = 0x10
	ImGuiOldColumnFlags_NoBorder                                 ImGuiOldColumnFlags          = 0x1
	ImGuiOldColumnFlags_NoForceWithinWindow                      ImGuiOldColumnFlags          = 0x8
	ImGuiOldColumnFlags_NoPreserveWidths                         ImGuiOldColumnFlags          = 0x4
	ImGuiOldColumnFlags_NoResize                                 ImGuiOldColumnFlags          = 0x2
	ImGuiOldColumnFlags_None                                     ImGuiOldColumnFlags          = 0x0
	ImGuiComboFlags_CustomPreview                                ImGuiComboFlagsPrivate_      = 0x100000
	ImGuiCond_Always                                             ImGuiCond                    = 0x1
	ImGuiCond_Appearing                                          ImGuiCond                    = 0x8
	ImGuiCond_FirstUseEver                                       ImGuiCond                    = 0x4
	ImGuiCond_None                                               ImGuiCond                    = 0x0
	ImGuiCond_Once                                               ImGuiCond                    = 0x2
	ImGuiDebugLogFlags_EventActiveId                             ImGuiDebugLogFlags           = 0x1
	ImGuiDebugLogFlags_EventClipper                              ImGuiDebugLogFlags           = 0x10
	ImGuiDebugLogFlags_EventDocking                              ImGuiDebugLogFlags           = 0x80
	ImGuiDebugLogFlags_EventFocus                                ImGuiDebugLogFlags           = 0x2
	ImGuiDebugLogFlags_EventIO                                   ImGuiDebugLogFlags           = 0x40
	ImGuiDebugLogFlags_EventMask_                                ImGuiDebugLogFlags           = 0x1ff
	ImGuiDebugLogFlags_EventNav                                  ImGuiDebugLogFlags           = 0x8
	ImGuiDebugLogFlags_EventPopup                                ImGuiDebugLogFlags           = 0x4
	ImGuiDebugLogFlags_EventSelection                            ImGuiDebugLogFlags           = 0x20
	ImGuiDebugLogFlags_EventViewport                             ImGuiDebugLogFlags           = 0x100
	ImGuiDebugLogFlags_None                                      ImGuiDebugLogFlags           = 0x0
	ImGuiDebugLogFlags_OutputToTTY                               ImGuiDebugLogFlags           = 0x400
	ImGuiTableRowFlags_Headers                                   ImGuiTableRowFlags           = 0x1
	ImGuiTableRowFlags_None                                      ImGuiTableRowFlags           = 0x0
	ImGuiWindowFlags_AlwaysAutoResize                            ImGuiWindowFlags             = 0x40
	ImGuiWindowFlags_AlwaysHorizontalScrollbar                   ImGuiWindowFlags             = 0x8000
	ImGuiWindowFlags_AlwaysUseWindowPadding                      ImGuiWindowFlags             = 0x10000
	ImGuiWindowFlags_AlwaysVerticalScrollbar                     ImGuiWindowFlags             = 0x4000
	ImGuiWindowFlags_ChildMenu                                   ImGuiWindowFlags             = 0x10000000
	ImGuiWindowFlags_ChildWindow                                 ImGuiWindowFlags             = 0x1000000
	ImGuiWindowFlags_DockNodeHost                                ImGuiWindowFlags             = 0x20000000
	ImGuiWindowFlags_HorizontalScrollbar                         ImGuiWindowFlags             = 0x800
	ImGuiWindowFlags_MenuBar                                     ImGuiWindowFlags             = 0x400
	ImGuiWindowFlags_Modal                                       ImGuiWindowFlags             = 0x8000000
	ImGuiWindowFlags_NavFlattened                                ImGuiWindowFlags             = 0x800000
	ImGuiWindowFlags_NoBackground                                ImGuiWindowFlags             = 0x80
	ImGuiWindowFlags_NoBringToFrontOnFocus                       ImGuiWindowFlags             = 0x2000
	ImGuiWindowFlags_NoCollapse                                  ImGuiWindowFlags             = 0x20
	ImGuiWindowFlags_NoDecoration                                ImGuiWindowFlags             = 0x2b
	ImGuiWindowFlags_NoDocking                                   ImGuiWindowFlags             = 0x200000
	ImGuiWindowFlags_NoFocusOnAppearing                          ImGuiWindowFlags             = 0x1000
	ImGuiWindowFlags_NoInputs                                    ImGuiWindowFlags             = 0xc0200
	ImGuiWindowFlags_NoMouseInputs                               ImGuiWindowFlags             = 0x200
	ImGuiWindowFlags_NoMove                                      ImGuiWindowFlags             = 0x4
	ImGuiWindowFlags_NoNav                                       ImGuiWindowFlags             = 0xc0000
	ImGuiWindowFlags_NoNavFocus                                  ImGuiWindowFlags             = 0x80000
	ImGuiWindowFlags_NoNavInputs                                 ImGuiWindowFlags             = 0x40000
	ImGuiWindowFlags_NoResize                                    ImGuiWindowFlags             = 0x2
	ImGuiWindowFlags_NoSavedSettings                             ImGuiWindowFlags             = 0x100
	ImGuiWindowFlags_NoScrollWithMouse                           ImGuiWindowFlags             = 0x10
	ImGuiWindowFlags_NoScrollbar                                 ImGuiWindowFlags             = 0x8
	ImGuiWindowFlags_NoTitleBar                                  ImGuiWindowFlags             = 0x1
	ImGuiWindowFlags_None                                        ImGuiWindowFlags             = 0x0
	ImGuiWindowFlags_Popup                                       ImGuiWindowFlags             = 0x4000000
	ImGuiWindowFlags_Tooltip                                     ImGuiWindowFlags             = 0x2000000
	ImGuiWindowFlags_UnsavedDocument                             ImGuiWindowFlags             = 0x100000
	ImDrawListFlags_AllowVtxOffset                               ImDrawListFlags              = 0x8
	ImDrawListFlags_AntiAliasedFill                              ImDrawListFlags              = 0x4
	ImDrawListFlags_AntiAliasedLines                             ImDrawListFlags              = 0x1
	ImDrawListFlags_AntiAliasedLinesUseTex                       ImDrawListFlags              = 0x2
	ImDrawListFlags_None                                         ImDrawListFlags              = 0x0
	ImGuiConfigFlags_DockingEnable                               ImGuiConfigFlags             = 0x40
	ImGuiConfigFlags_DpiEnableScaleFonts                         ImGuiConfigFlags             = 0x8000
	ImGuiConfigFlags_DpiEnableScaleViewports                     ImGuiConfigFlags             = 0x4000
	ImGuiConfigFlags_IsSRGB                                      ImGuiConfigFlags             = 0x100000
	ImGuiConfigFlags_IsTouchScreen                               ImGuiConfigFlags             = 0x200000
	ImGuiConfigFlags_NavEnableGamepad                            ImGuiConfigFlags             = 0x2
	ImGuiConfigFlags_NavEnableKeyboard                           ImGuiConfigFlags             = 0x1
	ImGuiConfigFlags_NavEnableSetMousePos                        ImGuiConfigFlags             = 0x4
	ImGuiConfigFlags_NavNoCaptureKeyboard                        ImGuiConfigFlags             = 0x8
	ImGuiConfigFlags_NoMouse                                     ImGuiConfigFlags             = 0x10
	ImGuiConfigFlags_NoMouseCursorChange                         ImGuiConfigFlags             = 0x20
	ImGuiConfigFlags_None                                        ImGuiConfigFlags             = 0x0
	ImGuiConfigFlags_ViewportsEnable                             ImGuiConfigFlags             = 0x400
	ImGuiMouseSource_COUNT                                       ImGuiMouseSource             = 0x3
	ImGuiMouseSource_Mouse                                       ImGuiMouseSource             = 0x0
	ImGuiMouseSource_Pen                                         ImGuiMouseSource             = 0x2
	ImGuiMouseSource_TouchScreen                                 ImGuiMouseSource             = 0x1
	ImGuiSelectableFlags_AllowDoubleClick                        ImGuiSelectableFlags         = 0x4
	ImGuiSelectableFlags_AllowOverlap                            ImGuiSelectableFlags         = 0x10
	ImGuiSelectableFlags_Disabled                                ImGuiSelectableFlags         = 0x8
	ImGuiSelectableFlags_DontClosePopups                         ImGuiSelectableFlags         = 0x1
	ImGuiSelectableFlags_None                                    ImGuiSelectableFlags         = 0x0
	ImGuiSelectableFlags_SpanAllColumns                          ImGuiSelectableFlags         = 0x2
	ImGuiSliderFlags_AlwaysClamp                                 ImGuiSliderFlags             = 0x10
	ImGuiSliderFlags_InvalidMask_                                ImGuiSliderFlags             = 0x7000000f
	ImGuiSliderFlags_Logarithmic                                 ImGuiSliderFlags             = 0x20
	ImGuiSliderFlags_NoInput                                     ImGuiSliderFlags             = 0x80
	ImGuiSliderFlags_NoRoundToFormat                             ImGuiSliderFlags             = 0x40
	ImGuiSliderFlags_None                                        ImGuiSliderFlags             = 0x0
	ImGuiTabItemFlags_Button                                     ImGuiTabItemFlagsPrivate_    = 0x200000
	ImGuiTabItemFlags_NoCloseButton                              ImGuiTabItemFlagsPrivate_    = 0x100000
	ImGuiTabItemFlags_Preview                                    ImGuiTabItemFlagsPrivate_    = 0x800000
	ImGuiTabItemFlags_SectionMask_                               ImGuiTabItemFlagsPrivate_    = 0xc0
	ImGuiTabItemFlags_Unsorted                                   ImGuiTabItemFlagsPrivate_    = 0x400000
	ImGuiTabItemFlags_Leading                                    ImGuiTabItemFlags            = 0x40
	ImGuiTabItemFlags_NoCloseWithMiddleMouseButton               ImGuiTabItemFlags            = 0x4
	ImGuiTabItemFlags_NoPushId                                   ImGuiTabItemFlags            = 0x8
	ImGuiTabItemFlags_NoReorder                                  ImGuiTabItemFlags            = 0x20
	ImGuiTabItemFlags_NoTooltip                                  ImGuiTabItemFlags            = 0x10
	ImGuiTabItemFlags_None                                       ImGuiTabItemFlags            = 0x0
	ImGuiTabItemFlags_SetSelected                                ImGuiTabItemFlags            = 0x2
	ImGuiTabItemFlags_Trailing                                   ImGuiTabItemFlags            = 0x80
	ImGuiTabItemFlags_UnsavedDocument                            ImGuiTabItemFlags            = 0x1
	ImFontAtlasFlags_NoBakedLines                                ImFontAtlasFlags             = 0x4
	ImFontAtlasFlags_NoMouseCursors                              ImFontAtlasFlags             = 0x2
	ImFontAtlasFlags_NoPowerOfTwoHeight                          ImFontAtlasFlags             = 0x1
	ImFontAtlasFlags_None                                        ImFontAtlasFlags             = 0x0
	ImGuiContextHookType_EndFramePost                            ImGuiContextHookType         = 0x3
	ImGuiContextHookType_EndFramePre                             ImGuiContextHookType         = 0x2
	ImGuiContextHookType_NewFramePost                            ImGuiContextHookType         = 0x1
	ImGuiContextHookType_NewFramePre                             ImGuiContextHookType         = 0x0
	ImGuiContextHookType_PendingRemoval_                         ImGuiContextHookType         = 0x7
	ImGuiContextHookType_RenderPost                              ImGuiContextHookType         = 0x5
	ImGuiContextHookType_RenderPre                               ImGuiContextHookType         = 0x4
	ImGuiContextHookType_Shutdown                                ImGuiContextHookType         = 0x6
	ImGuiDataType_COUNT                                          ImGuiDataType                = 0xa
	ImGuiDataType_Double                                         ImGuiDataType                = 0x9
	ImGuiDataType_Float                                          ImGuiDataType                = 0x8
	ImGuiDataType_S16                                            ImGuiDataType                = 0x2
	ImGuiDataType_S32                                            ImGuiDataType                = 0x4
	ImGuiDataType_S64                                            ImGuiDataType                = 0x6
	ImGuiDataType_S8                                             ImGuiDataType                = 0x0
	ImGuiDataType_U16                                            ImGuiDataType                = 0x3
	ImGuiDataType_U32                                            ImGuiDataType                = 0x5
	ImGuiDataType_U64                                            ImGuiDataType                = 0x7
	ImGuiDataType_U8                                             ImGuiDataType                = 0x1
	ImGuiTextFlags_NoWidthForLargeClippedText                    ImGuiTextFlags               = 0x1
	ImGuiTextFlags_None                                          ImGuiTextFlags               = 0x0
	ImGuiTooltipFlags_None                                       ImGuiTooltipFlags            = 0x0
	ImGuiTooltipFlags_OverridePrevious                           ImGuiTooltipFlags            = 0x2
	ImGuiComboFlags_HeightLarge                                  ImGuiComboFlags              = 0x8
	ImGuiComboFlags_HeightLargest                                ImGuiComboFlags              = 0x10
	ImGuiComboFlags_HeightMask_                                  ImGuiComboFlags              = 0x1e
	ImGuiComboFlags_HeightRegular                                ImGuiComboFlags              = 0x4
	ImGuiComboFlags_HeightSmall                                  ImGuiComboFlags              = 0x2
	ImGuiComboFlags_NoArrowButton                                ImGuiComboFlags              = 0x20
	ImGuiComboFlags_NoPreview                                    ImGuiComboFlags              = 0x40
	ImGuiComboFlags_None                                         ImGuiComboFlags              = 0x0
	ImGuiComboFlags_PopupAlignLeft                               ImGuiComboFlags              = 0x1
	ImGuiDockNodeFlags_CentralNode                               ImGuiDockNodeFlagsPrivate_   = 0x800
	ImGuiDockNodeFlags_DockSpace                                 ImGuiDockNodeFlagsPrivate_   = 0x400
	ImGuiDockNodeFlags_HiddenTabBar                              ImGuiDockNodeFlagsPrivate_   = 0x2000
	ImGuiDockNodeFlags_LocalFlagsMask_                           ImGuiDockNodeFlagsPrivate_   = 0xc1fc70
	ImGuiDockNodeFlags_LocalFlagsTransferMask_                   ImGuiDockNodeFlagsPrivate_   = 0xc1f870
	ImGuiDockNodeFlags_NoCloseButton                             ImGuiDockNodeFlagsPrivate_   = 0x8000
	ImGuiDockNodeFlags_NoDocking                                 ImGuiDockNodeFlagsPrivate_   = 0x10000
	ImGuiDockNodeFlags_NoDockingOverEmpty                        ImGuiDockNodeFlagsPrivate_   = 0x200000
	ImGuiDockNodeFlags_NoDockingOverMe                           ImGuiDockNodeFlagsPrivate_   = 0x80000
	ImGuiDockNodeFlags_NoDockingOverOther                        ImGuiDockNodeFlagsPrivate_   = 0x100000
	ImGuiDockNodeFlags_NoDockingSplitMe                          ImGuiDockNodeFlagsPrivate_   = 0x20000
	ImGuiDockNodeFlags_NoDockingSplitOther                       ImGuiDockNodeFlagsPrivate_   = 0x40000
	ImGuiDockNodeFlags_NoResizeFlagsMask_                        ImGuiDockNodeFlagsPrivate_   = 0xc00020
	ImGuiDockNodeFlags_NoResizeX                                 ImGuiDockNodeFlagsPrivate_   = 0x400000
	ImGuiDockNodeFlags_NoResizeY                                 ImGuiDockNodeFlagsPrivate_   = 0x800000
	ImGuiDockNodeFlags_NoTabBar                                  ImGuiDockNodeFlagsPrivate_   = 0x1000
	ImGuiDockNodeFlags_NoWindowMenuButton                        ImGuiDockNodeFlagsPrivate_   = 0x4000
	ImGuiDockNodeFlags_SavedFlagsMask_                           ImGuiDockNodeFlagsPrivate_   = 0xc1fc20
	ImGuiDockNodeFlags_SharedFlagsInheritMask_                   ImGuiDockNodeFlagsPrivate_   = -0x1
	ImGuiSeparatorFlags_Horizontal                               ImGuiSeparatorFlags          = 0x1
	ImGuiSeparatorFlags_None                                     ImGuiSeparatorFlags          = 0x0
	ImGuiSeparatorFlags_SpanAllColumns                           ImGuiSeparatorFlags          = 0x4
	ImGuiSeparatorFlags_Vertical                                 ImGuiSeparatorFlags          = 0x2
	ImGuiDataType_ID                                             ImGuiDataTypePrivate_        = 0xd
	ImGuiDataType_Pointer                                        ImGuiDataTypePrivate_        = 0xc
	ImGuiDataType_String                                         ImGuiDataTypePrivate_        = 0xb
	ImGuiInputEventType_COUNT                                    ImGuiInputEventType          = 0x8
	ImGuiInputEventType_Focus                                    ImGuiInputEventType          = 0x7
	ImGuiInputEventType_Key                                      ImGuiInputEventType          = 0x5
	ImGuiInputEventType_MouseButton                              ImGuiInputEventType          = 0x3
	ImGuiInputEventType_MousePos                                 ImGuiInputEventType          = 0x1
	ImGuiInputEventType_MouseViewport                            ImGuiInputEventType          = 0x4
	ImGuiInputEventType_MouseWheel                               ImGuiInputEventType          = 0x2
	ImGuiInputEventType_None                                     ImGuiInputEventType          = 0x0
	ImGuiInputEventType_Text                                     ImGuiInputEventType          = 0x6
	ImGuiMouseCursor_Arrow                                       ImGuiMouseCursor             = 0x0
	ImGuiMouseCursor_COUNT                                       ImGuiMouseCursor             = 0x9
	ImGuiMouseCursor_Hand                                        ImGuiMouseCursor             = 0x7
	ImGuiMouseCursor_None                                        ImGuiMouseCursor             = -0x1
	ImGuiMouseCursor_NotAllowed                                  ImGuiMouseCursor             = 0x8
	ImGuiMouseCursor_ResizeAll                                   ImGuiMouseCursor             = 0x2
	ImGuiMouseCursor_ResizeEW                                    ImGuiMouseCursor             = 0x4
	ImGuiMouseCursor_ResizeNESW                                  ImGuiMouseCursor             = 0x5
	ImGuiMouseCursor_ResizeNS                                    ImGuiMouseCursor             = 0x3
	ImGuiMouseCursor_ResizeNWSE                                  ImGuiMouseCursor             = 0x6
	ImGuiMouseCursor_TextInput                                   ImGuiMouseCursor             = 0x1
	ImGuiNavLayer_COUNT                                          ImGuiNavLayer                = 0x2
	ImGuiNavLayer_Main                                           ImGuiNavLayer                = 0x0
	ImGuiNavLayer_Menu                                           ImGuiNavLayer                = 0x1
	ImGuiSliderFlags_ReadOnly                                    ImGuiSliderFlagsPrivate_     = 0x200000
	ImGuiSliderFlags_Vertical                                    ImGuiSliderFlagsPrivate_     = 0x100000
	ImGuiStyleVar_Alpha                                          ImGuiStyleVar                = 0x0
	ImGuiStyleVar_ButtonTextAlign                                ImGuiStyleVar                = 0x17
	ImGuiStyleVar_COUNT                                          ImGuiStyleVar                = 0x1c
	ImGuiStyleVar_CellPadding                                    ImGuiStyleVar                = 0x11
	ImGuiStyleVar_ChildBorderSize                                ImGuiStyleVar                = 0x8
	ImGuiStyleVar_ChildRounding                                  ImGuiStyleVar                = 0x7
	ImGuiStyleVar_DisabledAlpha                                  ImGuiStyleVar                = 0x1
	ImGuiStyleVar_FrameBorderSize                                ImGuiStyleVar                = 0xd
	ImGuiStyleVar_FramePadding                                   ImGuiStyleVar                = 0xb
	ImGuiStyleVar_FrameRounding                                  ImGuiStyleVar                = 0xc
	ImGuiStyleVar_GrabMinSize                                    ImGuiStyleVar                = 0x14
	ImGuiStyleVar_GrabRounding                                   ImGuiStyleVar                = 0x15
	ImGuiStyleVar_IndentSpacing                                  ImGuiStyleVar                = 0x10
	ImGuiStyleVar_ItemInnerSpacing                               ImGuiStyleVar                = 0xf
	ImGuiStyleVar_ItemSpacing                                    ImGuiStyleVar                = 0xe
	ImGuiStyleVar_PopupBorderSize                                ImGuiStyleVar                = 0xa
	ImGuiStyleVar_PopupRounding                                  ImGuiStyleVar                = 0x9
	ImGuiStyleVar_ScrollbarRounding                              ImGuiStyleVar                = 0x13
	ImGuiStyleVar_ScrollbarSize                                  ImGuiStyleVar                = 0x12
	ImGuiStyleVar_SelectableTextAlign                            ImGuiStyleVar                = 0x18
	ImGuiStyleVar_SeparatorTextAlign                             ImGuiStyleVar                = 0x1a
	ImGuiStyleVar_SeparatorTextBorderSize                        ImGuiStyleVar                = 0x19
	ImGuiStyleVar_SeparatorTextPadding                           ImGuiStyleVar                = 0x1b
	ImGuiStyleVar_TabRounding                                    ImGuiStyleVar                = 0x16
	ImGuiStyleVar_WindowBorderSize                               ImGuiStyleVar                = 0x4
	ImGuiStyleVar_WindowMinSize                                  ImGuiStyleVar                = 0x5
	ImGuiStyleVar_WindowPadding                                  ImGuiStyleVar                = 0x2
	ImGuiStyleVar_WindowRounding                                 ImGuiStyleVar                = 0x3
	ImGuiStyleVar_WindowTitleAlign                               ImGuiStyleVar                = 0x6
	ImGuiViewportFlags_CanHostOtherWindows                       ImGuiViewportFlags           = 0x800
	ImGuiViewportFlags_IsFocused                                 ImGuiViewportFlags           = 0x2000
	ImGuiViewportFlags_IsMinimized                               ImGuiViewportFlags           = 0x1000
	ImGuiViewportFlags_IsPlatformMonitor                         ImGuiViewportFlags           = 0x2
	ImGuiViewportFlags_IsPlatformWindow                          ImGuiViewportFlags           = 0x1
	ImGuiViewportFlags_NoAutoMerge                               ImGuiViewportFlags           = 0x200
	ImGuiViewportFlags_NoDecoration                              ImGuiViewportFlags           = 0x8
	ImGuiViewportFlags_NoFocusOnAppearing                        ImGuiViewportFlags           = 0x20
	ImGuiViewportFlags_NoFocusOnClick                            ImGuiViewportFlags           = 0x40
	ImGuiViewportFlags_NoInputs                                  ImGuiViewportFlags           = 0x80
	ImGuiViewportFlags_NoRendererClear                           ImGuiViewportFlags           = 0x100
	ImGuiViewportFlags_NoTaskBarIcon                             ImGuiViewportFlags           = 0x10
	ImGuiViewportFlags_None                                      ImGuiViewportFlags           = 0x0
	ImGuiViewportFlags_OwnedByApp                                ImGuiViewportFlags           = 0x4
	ImGuiViewportFlags_TopMost                                   ImGuiViewportFlags           = 0x400
	ImGuiDataAuthority_Auto                                      ImGuiDataAuthority           = 0x0
	ImGuiDataAuthority_DockNode                                  ImGuiDataAuthority           = 0x1
	ImGuiDataAuthority_Window                                    ImGuiDataAuthority           = 0x2
	ImGuiInputFlags_CondActive                                   ImGuiInputFlags              = 0x20
	ImGuiInputFlags_CondDefault_                                 ImGuiInputFlags              = 0x30
	ImGuiInputFlags_CondHovered                                  ImGuiInputFlags              = 0x10
	ImGuiInputFlags_CondMask_                                    ImGuiInputFlags              = 0x30
	ImGuiInputFlags_LockThisFrame                                ImGuiInputFlags              = 0x40
	ImGuiInputFlags_LockUntilRelease                             ImGuiInputFlags              = 0x80
	ImGuiInputFlags_None                                         ImGuiInputFlags              = 0x0
	ImGuiInputFlags_Repeat                                       ImGuiInputFlags              = 0x1
	ImGuiInputFlags_RepeatRateDefault                            ImGuiInputFlags              = 0x2
	ImGuiInputFlags_RepeatRateMask_                              ImGuiInputFlags              = 0xe
	ImGuiInputFlags_RepeatRateNavMove                            ImGuiInputFlags              = 0x4
	ImGuiInputFlags_RepeatRateNavTweak                           ImGuiInputFlags              = 0x8
	ImGuiInputFlags_RouteAlways                                  ImGuiInputFlags              = 0x1000
	ImGuiInputFlags_RouteExtraMask_                              ImGuiInputFlags              = 0x3000
	ImGuiInputFlags_RouteFocused                                 ImGuiInputFlags              = 0x100
	ImGuiInputFlags_RouteGlobal                                  ImGuiInputFlags              = 0x400
	ImGuiInputFlags_RouteGlobalHigh                              ImGuiInputFlags              = 0x800
	ImGuiInputFlags_RouteGlobalLow                               ImGuiInputFlags              = 0x200
	ImGuiInputFlags_RouteMask_                                   ImGuiInputFlags              = 0xf00
	ImGuiInputFlags_RouteUnlessBgFocused                         ImGuiInputFlags              = 0x2000
	ImGuiInputFlags_SupportedByIsKeyPressed                      ImGuiInputFlags              = 0xf
	ImGuiInputFlags_SupportedBySetItemKeyOwner                   ImGuiInputFlags              = 0xf0
	ImGuiInputFlags_SupportedBySetKeyOwner                       ImGuiInputFlags              = 0xc0
	ImGuiInputFlags_SupportedByShortcut                          ImGuiInputFlags              = 0x3f0f
	ImGuiInputTextFlags_MergedItem                               ImGuiInputTextFlagsPrivate_  = 0x10000000
	ImGuiInputTextFlags_Multiline                                ImGuiInputTextFlagsPrivate_  = 0x4000000
	ImGuiInputTextFlags_NoMarkEdited                             ImGuiInputTextFlagsPrivate_  = 0x8000000
	ImGuiNavMoveFlags_Activate                                   ImGuiNavMoveFlags            = 0x800
	ImGuiNavMoveFlags_AllowCurrentNavId                          ImGuiNavMoveFlags            = 0x10
	ImGuiNavMoveFlags_AlsoScoreVisibleSet                        ImGuiNavMoveFlags            = 0x20
	ImGuiNavMoveFlags_DebugNoResult                              ImGuiNavMoveFlags            = 0x100
	ImGuiNavMoveFlags_FocusApi                                   ImGuiNavMoveFlags            = 0x200
	ImGuiNavMoveFlags_Forwarded                                  ImGuiNavMoveFlags            = 0x80
	ImGuiNavMoveFlags_LoopX                                      ImGuiNavMoveFlags            = 0x1
	ImGuiNavMoveFlags_LoopY                                      ImGuiNavMoveFlags            = 0x2
	ImGuiNavMoveFlags_NoSelect                                   ImGuiNavMoveFlags            = 0x1000
	ImGuiNavMoveFlags_NoSetNavHighlight                          ImGuiNavMoveFlags            = 0x2000
	ImGuiNavMoveFlags_None                                       ImGuiNavMoveFlags            = 0x0
	ImGuiNavMoveFlags_ScrollToEdgeY                              ImGuiNavMoveFlags            = 0x40
	ImGuiNavMoveFlags_Tabbing                                    ImGuiNavMoveFlags            = 0x400
	ImGuiNavMoveFlags_WrapMask_                                  ImGuiNavMoveFlags            = 0xf
	ImGuiNavMoveFlags_WrapX                                      ImGuiNavMoveFlags            = 0x4
	ImGuiNavMoveFlags_WrapY                                      ImGuiNavMoveFlags            = 0x8
	ImGuiScrollFlags_AlwaysCenterX                               ImGuiScrollFlags             = 0x10
	ImGuiScrollFlags_AlwaysCenterY                               ImGuiScrollFlags             = 0x20
	ImGuiScrollFlags_KeepVisibleCenterX                          ImGuiScrollFlags             = 0x4
	ImGuiScrollFlags_KeepVisibleCenterY                          ImGuiScrollFlags             = 0x8
	ImGuiScrollFlags_KeepVisibleEdgeX                            ImGuiScrollFlags             = 0x1
	ImGuiScrollFlags_KeepVisibleEdgeY                            ImGuiScrollFlags             = 0x2
	ImGuiScrollFlags_MaskX_                                      ImGuiScrollFlags             = 0x15
	ImGuiScrollFlags_MaskY_                                      ImGuiScrollFlags             = 0x2a
	ImGuiScrollFlags_NoScrollParent                              ImGuiScrollFlags             = 0x40
	ImGuiScrollFlags_None                                        ImGuiScrollFlags             = 0x0
	ImGuiTabBarFlags_DockNode                                    ImGuiTabBarFlagsPrivate_     = 0x100000
	ImGuiTabBarFlags_IsFocused                                   ImGuiTabBarFlagsPrivate_     = 0x200000
	ImGuiTabBarFlags_SaveSettings                                ImGuiTabBarFlagsPrivate_     = 0x400000
	ImGuiHoveredFlags_AllowWhenBlockedByActiveItem               ImGuiHoveredFlags            = 0x80
	ImGuiHoveredFlags_AllowWhenBlockedByPopup                    ImGuiHoveredFlags            = 0x20
	ImGuiHoveredFlags_AllowWhenDisabled                          ImGuiHoveredFlags            = 0x400
	ImGuiHoveredFlags_AllowWhenOverlapped                        ImGuiHoveredFlags            = 0x300
	ImGuiHoveredFlags_AllowWhenOverlappedByItem                  ImGuiHoveredFlags            = 0x100
	ImGuiHoveredFlags_AllowWhenOverlappedByWindow                ImGuiHoveredFlags            = 0x200
	ImGuiHoveredFlags_AnyWindow                                  ImGuiHoveredFlags            = 0x4
	ImGuiHoveredFlags_ChildWindows                               ImGuiHoveredFlags            = 0x1
	ImGuiHoveredFlags_DelayNone                                  ImGuiHoveredFlags            = 0x2000
	ImGuiHoveredFlags_DelayNormal                                ImGuiHoveredFlags            = 0x8000
	ImGuiHoveredFlags_DelayShort                                 ImGuiHoveredFlags            = 0x4000
	ImGuiHoveredFlags_DockHierarchy                              ImGuiHoveredFlags            = 0x10
	ImGuiHoveredFlags_ForTooltip                                 ImGuiHoveredFlags            = 0x800
	ImGuiHoveredFlags_NoNavOverride                              ImGuiHoveredFlags            = 0x800
	ImGuiHoveredFlags_NoPopupHierarchy                           ImGuiHoveredFlags            = 0x8
	ImGuiHoveredFlags_NoSharedDelay                              ImGuiHoveredFlags            = 0x10000
	ImGuiHoveredFlags_None                                       ImGuiHoveredFlags            = 0x0
	ImGuiHoveredFlags_RectOnly                                   ImGuiHoveredFlags            = 0x3a0
	ImGuiHoveredFlags_RootAndChildWindows                        ImGuiHoveredFlags            = 0x3
	ImGuiHoveredFlags_RootWindow                                 ImGuiHoveredFlags            = 0x2
	ImGuiHoveredFlags_Stationary                                 ImGuiHoveredFlags            = 0x1000
)

// type ImGuiInputFlags int32
// type ImGuiOldColumnFlags int32
// type ImGuiStyleVar int32

// struct define
type (
	ImSpan_ImGuiTableColumn struct {
		Data    *ImGuiTableColumn
		DataEnd *ImGuiTableColumn
	}
	ImVector_ImGuiSettingsHandler struct {
		Size     int32
		Capacity int32
		Data     *ImGuiSettingsHandler
	}
	ImVector_ImGuiTableTempData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTableTempData
	}
	ImBitArrayForNamedKeys struct {
		Storage [5]uint32
	}
	ImChunkStream_ImGuiWindowSettings struct {
		Buf ImVector_char
	}
	ImFileHandle              *C.struct__iobuf
	ImPoolIdx                 int32
	ImVector_ImGuiContextHook struct {
		Size     int32
		Capacity int32
		Data     *ImGuiContextHook
	}
	ImGuiInputTextCallback *[0]byte

	// ImGuiPopupFlags int32
	// ImGuiSeparatorFlags int32

	ImVector_ImDrawCmd struct {
		Size     int32
		Capacity int32
		Data     *ImDrawCmd
	}

	// ImFontAtlasFlags int32
	// ImGuiDataAuthority int32

	ImU16          uint16
	ImVector_ImU32 struct {
		Size     int32
		Capacity int32
		Data     *uint32
	}
	ImChunkStream_ImGuiTableSettings struct {
		Buf ImVector_char
	}

	// ImGuiMouseCursor int32
	// ImGuiNavHighlightFlags int32
	// ImGuiNextItemDataFlags int32
	ImU8               uint8
	ImVector_ImDrawIdx struct {
		Size     int32
		Capacity int32
		Data     *uint16
	}
	ImVector_unsigned_char struct {
		Size     int32
		Capacity int32
		Data     *uint8
	}

	// ImGuiActivateFlags int32
	// ImGuiButtonFlags int32
	// ImGuiDockNodeFlags int32

	ImGuiID                  uint32
	ImVector_ImGuiOldColumns struct {
		Size     int32
		Capacity int32
		Data     *ImGuiOldColumns
	}
	ImVector_float struct {
		Size     int32
		Capacity int32
		Data     *float32
	}
	StbTexteditRow struct {
		X0               float32
		X1               float32
		Baseline_y_delta float32
		Ymin             float32
		Ymax             float32
		Num_chars        int32
	}
	ImVector_ImGuiListClipperData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiListClipperData
	}
	StbUndoRecord struct {
		Where         int32
		Insert_length int32
		Delete_length int32
		Char_storage  int32
	}
	ImBitArrayPtr            *uint32
	ImGuiTableDrawChannelIdx uint16
	ImVector_ImGuiItemFlags  struct {
		Size     int32
		Capacity int32
		Data     *int32
	}
	ImVector_ImGuiWindowPtr struct {
		Size     int32
		Capacity int32
		Data     **ImGuiWindow
	}
	ImWchar32 uint32

	// ImDrawFlags int32

	ImGuiSizeCallback  *[0]byte
	ImVector_ImFontPtr struct {
		Size     int32
		Capacity int32
		Data     **ImFont
	}

	// ImGuiCond int32
	ImU64                  uint64
	ImVector_ImGuiColorMod struct {
		Size     int32
		Capacity int32
		Data     *ImGuiColorMod
	}
	ImVector_ImGuiID struct {
		Size     int32
		Capacity int32
		Data     *uint32
	}

	// ImGuiFocusedFlags int32
	// ImGuiTooltipFlags int32

	ImVector_ImGuiOldColumnData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiOldColumnData
	}
	ImVector_ImGuiPlatformMonitor struct {
		Size     int32
		Capacity int32
		Data     *ImGuiPlatformMonitor
	}
	ImVector_ImGuiTableInstanceData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTableInstanceData
	}
	ImDrawIdx uint16

	// ImGuiItemFlags int32
	// ImGuiTableRowFlags int32

	ImPool_ImGuiTable struct {
		Buf        ImVector_ImGuiTable
		Map        ImGuiStorage
		FreeIdx    int32
		AliveCount int32
	}
	ImSpan_ImGuiTableColumnIdx struct {
		Data    *int16
		DataEnd *int16
	}
	ImVector_ImDrawVert struct {
		Size     int32
		Capacity int32
		Data     *ImDrawVert
	}

	// ImGuiItemStatusFlags int32
	ImGuiMemAllocFunc *[0]byte

	// ImGuiSelectableFlags int32
	// ImGuiSortDirection int32
	// ImGuiViewportFlags int32

	ImVector_const_charPtr struct {
		Size     int32
		Capacity int32
		Data     **int8
	}
	ImWchar16             uint16
	ImGuiErrorLogCallback *[0]byte
	ImGuiMemFreeFunc      *[0]byte
	ImVector_ImFontConfig struct {
		Size     int32
		Capacity int32
		Data     *ImFontConfig
	}
	ImVector_ImGuiPopupData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiPopupData
	}
	ImVector_ImGuiTableColumnSortSpecs struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTableColumnSortSpecs
	}
	ImVector_ImGuiViewportPtr struct {
		Size     int32
		Capacity int32
		Data     **ImGuiViewport
	}
	ImVector_ImWchar struct {
		Size     int32
		Capacity int32
		Data     *uint16
	}

	// ImGuiBackendFlags int32
	// ImGuiHoveredFlags int32
	ImVector_ImGuiKeyRoutingData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiKeyRoutingData
	}
	ImVector_ImGuiViewportPPtr struct {
		Size     int32
		Capacity int32
		Data     **ImGuiViewportP
	}
	ImVector_ImTextureID struct {
		Size     int32
		Capacity int32
		Data     **byte
	}
	StbUndoState struct {
		Undo_rec        [99]StbUndoRecord
		Undo_char       [999]uint16
		Undo_point      int16
		Redo_point      int16
		Undo_char_point int32
		Redo_char_point int32
	}
	ImDrawCallback           *[0]byte
	ImGuiContextHookCallback *[0]byte

	// ImGuiLayoutType int32
	ImVector_ImGuiDockRequest struct {
		Size     int32
		Capacity int32
		// Data     *ImGuiDockRequest
		Data unsafe.Pointer
	}
	ImVector_ImGuiTextRange struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTextRange
	}

	// ImGuiColorEditFlags int32
	// ImGuiDataType int32
	// ImGuiDebugLogFlags int32

	ImVector_ImVec4 struct {
		Size     int32
		Capacity int32
		Data     *ImVec4
	}
	ImGuiKeyRoutingIndex int16

	// ImGuiSliderFlags int32
	// ImGuiTabBarFlags int32
	// ImGuiTabItemFlags int32

	ImS64         int64
	ImVector_char struct {
		Size     int32
		Capacity int32
		Data     *int8
	}

	// ImGuiComboFlags int32
	// ImGuiFocusRequestFlags int32
	// ImGuiTableFlags int32

	ImPool_ImGuiTabBar struct {
		Buf        ImVector_ImGuiTabBar
		Map        ImGuiStorage
		FreeIdx    int32
		AliveCount int32
	}
	ImS16                 int16
	ImU32                 uint32
	ImVector_ImGuiTabItem struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTabItem
	}

	// ImGuiTableColumnFlags int32

	ImVector_ImGuiDockNodeSettings struct {
		Size     int32
		Capacity int32
		// Data     *ImGuiDockNodeSettings
		Data unsafe.Pointer
	}
	ImVector_ImGuiShrinkWidthItem struct {
		Size     int32
		Capacity int32
		Data     *ImGuiShrinkWidthItem
	}
	STB_TexteditState struct {
		Cursor                int32
		Select_start          int32
		Select_end            int32
		Insert_mode           uint8
		Row_count_per_page    int32
		Cursor_at_end_of_line uint8
		Initialized           uint8
		Has_preferred_x       uint8
		Single_line           uint8
		Padding1              uint8
		Padding2              uint8
		Padding3              uint8
		Preferred_x           float32
		Undostate             StbUndoState
	}

	ImS32                        int32
	ImVector_ImGuiStackLevelInfo struct {
		Size     int32
		Capacity int32
		Data     *ImGuiStackLevelInfo
	}
	ImGuiTableColumnIdx int16

	// ImGuiTextFlags int32
	// ImGuiTreeNodeFlags int32

	ImVector_ImDrawChannel struct {
		Size     int32
		Capacity int32
		Data     *ImDrawChannel
	}
	ImVector_ImFontGlyph struct {
		Size     int32
		Capacity int32
		Data     *ImFontGlyph
	}
	ImVector_ImGuiListClipperRange struct {
		Size     int32
		Capacity int32
		Data     *ImGuiListClipperRange
	}
	ImVector_int struct {
		Size     int32
		Capacity int32
		Data     *int32
	}

	// ImGuiConfigFlags int32
	// ImGuiInputTextFlags int32
	// ImGuiMouseButton int32
	// ImGuiNavMoveFlags int32
	// ImGuiNextWindowDataFlags int32
	// ImGuiTableBgTarget int32
	// ImGuiWindowFlags int32
	// ImGuiDragDropFlags int32

	ImVector_ImGuiWindowStackData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiWindowStackData
	}

	// ImGuiDir int32
	ImVector_ImFontAtlasCustomRect struct {
		Size     int32
		Capacity int32
		Data     *ImFontAtlasCustomRect
	}
	ImVector_ImGuiGroupData struct {
		Size     int32
		Capacity int32
		Data     *ImGuiGroupData
	}
	ImVector_ImVec2 struct {
		Size     int32
		Capacity int32
		Data     *ImVec2
	}
	ImWchar                  uint16
	ImTextureID              unsafe.Pointer
	ImVector_ImGuiPtrOrIndex struct {
		Size     int32
		Capacity int32
		Data     *ImGuiPtrOrIndex
	}

	// ImGuiCol int32
	// ImGuiScrollFlags int32

	ImVector_ImDrawListPtr struct {
		Size     int32
		Capacity int32
		Data     **ImDrawList
	}
	ImVector_ImGuiStyleMod struct {
		Size     int32
		Capacity int32
		Data     *ImGuiStyleMod
	}

	// ImDrawListFlags int32

	ImGuiKeyChord             int32
	ImS8                      int8
	ImSpan_ImGuiTableCellData struct {
		Data    *ImGuiTableCellData
		DataEnd *ImGuiTableCellData
	}
	ImVector_ImGuiInputEvent struct {
		Size     int32
		Capacity int32
		Data     *ImGuiInputEvent
	}
	ImVector_ImGuiStoragePair struct {
		Size     int32
		Capacity int32
		Data     *ImGuiStoragePair
	}

	ImDrawData struct {
		Valid            bool
		CmdListsCount    int32
		TotalIdxCount    int32
		TotalVtxCount    int32
		CmdLists         **ImDrawList
		DisplayPos       ImVec2
		DisplaySize      ImVec2
		FramebufferScale ImVec2
		OwnerViewport    *ImGuiViewport
	}
	ImFontAtlasCustomRect struct {
		Width         uint16
		Height        uint16
		X             uint16
		Y             uint16
		GlyphID       uint32
		GlyphAdvanceX float32
		GlyphOffset   ImVec2
		Font          *ImFont
	}
	ImGuiTableColumnSortSpecs struct {
		ColumnUserID uint32
		ColumnIndex  int16
		SortOrder    int16
		Pad_cgo_0    [4]byte
	}
	ImGuiTextBuffer struct {
		Buf ImVector_char
	}
	ImVector_ImGuiTable struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTable
	}
	ImDrawCmdHeader struct {
		ClipRect  ImVec4
		TextureId *byte
		VtxOffset uint32
		Pad_cgo_0 [4]byte
	}
	ImVector_ImGuiTabBar struct {
		Size     int32
		Capacity int32
		Data     *ImGuiTabBar
	}

	// ImGuiDockNodeSettings _cgopackage.Incomplete
	ImDrawChannel struct {
		X_CmdBuffer ImVector_ImDrawCmd
		X_IdxBuffer ImVector_ImDrawIdx
	}
	ImFontBuilderIO struct {
		Build *[0]byte
	}
	ImFontConfig struct {
		FontData             *byte
		FontDataSize         int32
		FontDataOwnedByAtlas bool
		FontNo               int32
		SizePixels           float32
		OversampleH          int32
		OversampleV          int32
		PixelSnapH           bool
		GlyphExtraSpacing    ImVec2
		GlyphOffset          ImVec2
		GlyphRanges          *uint16
		GlyphMinAdvanceX     float32
		GlyphMaxAdvanceX     float32
		MergeMode            bool
		FontBuilderFlags     uint32
		RasterizerMultiply   float32
		EllipsisChar         uint16
		Name                 [40]int8
		DstFont              *ImFont
	}
	ImGuiListClipperData struct {
		ListClipper     *ImGuiListClipper
		LossynessOffset float32
		StepNo          int32
		ItemsFrozen     int32
		Ranges          ImVector_ImGuiListClipperRange
	}
	ImGuiSizeCallbackData struct {
		UserData    *byte
		Pos         ImVec2
		CurrentSize ImVec2
		DesiredSize ImVec2
	}
	ImVec2ih struct {
		X int16
		Y int16
	}
	ImGuiInputEventMouseWheel struct {
		WheelX      float32
		WheelY      float32
		MouseSource uint32
	}
	ImGuiPayload struct {
		Data           *byte
		DataSize       int32
		SourceId       uint32
		SourceParentId uint32
		DataFrameCount int32
		DataType       [33]int8
		Preview        bool
		Delivery       bool
		Pad_cgo_0      [5]byte
	}
	ImGuiSettingsHandler struct {
		TypeName   *int8
		TypeHash   uint32
		ClearAllFn *[0]byte
		ReadInitFn *[0]byte
		ReadOpenFn *[0]byte
		ReadLineFn *[0]byte
		ApplyAllFn *[0]byte
		WriteAllFn *[0]byte
		UserData   *byte
	}

	// ImGuiTableColumnsSettings _cgopackage.Incomplete
	ImGuiTableInstanceData struct {
		TableInstanceID    uint32
		LastOuterHeight    float32
		LastFirstRowHeight float32
		LastFrozenHeight   float32
	}
	ImGuiListClipperRange struct {
		Min                 int32
		Max                 int32
		PosToIndexConvert   bool
		PosToIndexOffsetMin int8
		PosToIndexOffsetMax int8
		Pad_cgo_0           [1]byte
	}
	ImGuiInputEventMousePos struct {
		PosX        float32
		PosY        float32
		MouseSource uint32
	}
	ImGuiInputTextCallbackData struct {
		Ctx            *ImGuiContext
		EventFlag      int32
		Flags          int32
		UserData       *byte
		EventChar      uint16
		EventKey       uint32
		Buf            *int8
		BufTextLen     int32
		BufSize        int32
		BufDirty       bool
		CursorPos      int32
		SelectionStart int32
		SelectionEnd   int32
	}
	ImGuiInputEventText struct {
		Char uint32
	}
	ImGuiNextWindowData struct {
		Flags                int32
		PosCond              int32
		SizeCond             int32
		CollapsedCond        int32
		DockCond             int32
		PosVal               ImVec2
		PosPivotVal          ImVec2
		SizeVal              ImVec2
		ContentSizeVal       ImVec2
		ScrollVal            ImVec2
		PosUndock            bool
		CollapsedVal         bool
		SizeConstraintRect   ImRect
		SizeCallback         *[0]byte
		SizeCallbackUserData *byte
		BgAlphaVal           float32
		ViewportId           uint32
		DockId               uint32
		WindowClass          ImGuiWindowClass
		MenuBarOffsetMinVal  ImVec2
	}
	ImDrawDataBuilder struct {
		Layers [2]ImVector_ImDrawListPtr
	}
	ImDrawListSplitter struct {
		X_Current  int32
		X_Count    int32
		X_Channels ImVector_ImDrawChannel
	}
	ImFont struct {
		IndexAdvanceX       ImVector_float
		FallbackAdvanceX    float32
		FontSize            float32
		IndexLookup         ImVector_ImWchar
		Glyphs              ImVector_ImFontGlyph
		FallbackGlyph       *ImFontGlyph
		ContainerAtlas      *ImFontAtlas
		ConfigData          *ImFontConfig
		ConfigDataCount     int16
		FallbackChar        uint16
		EllipsisChar        uint16
		EllipsisCharCount   int16
		EllipsisWidth       float32
		EllipsisCharStep    float32
		DirtyLookupTables   bool
		Scale               float32
		Ascent              float32
		Descent             float32
		MetricsTotalSurface int32
		Used4kPagesMap      [2]uint8
		Pad_cgo_0           [2]byte
	}
	ImFontGlyphRangesBuilder struct {
		UsedChars ImVector_ImU32
	}
	ImGuiKeyRoutingTable struct {
		Index       [140]int16
		Entries     ImVector_ImGuiKeyRoutingData
		EntriesNext ImVector_ImGuiKeyRoutingData
	}
	ImGuiMetricsConfig struct {
		ShowDebugLog                 bool
		ShowStackTool                bool
		ShowWindowsRects             bool
		ShowWindowsBeginOrder        bool
		ShowTablesRects              bool
		ShowDrawCmdMesh              bool
		ShowDrawCmdBoundingBoxes     bool
		ShowAtlasTintedWithTextColor bool
		ShowDockingNodes             bool
		ShowWindowsRectsType         int32
		ShowTablesRectsType          int32
	}
	ImVec4 struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	ImGuiDataTypeTempStorage struct {
		Data [8]uint8
	}
	ImGuiInputEventKey struct {
		Key         uint32
		Down        bool
		AnalogValue float32
	}
	ImGuiOldColumns struct {
		ID                       uint32
		Flags                    int32
		IsFirstFrame             bool
		IsBeingResized           bool
		Current                  int32
		Count                    int32
		OffMinX                  float32
		OffMaxX                  float32
		LineMinY                 float32
		LineMaxY                 float32
		HostCursorPosY           float32
		HostCursorMaxPosX        float32
		HostInitialClipRect      ImRect
		HostBackupClipRect       ImRect
		HostBackupParentWorkRect ImRect
		Columns                  ImVector_ImGuiOldColumnData
		Splitter                 ImDrawListSplitter
	}
	ImGuiTextIndex struct {
		LineOffsets ImVector_int
		EndOffset   int32
		Pad_cgo_0   [4]byte
	}
	ImVec1 struct {
		X float32
	}
	ImDrawList struct {
		CmdBuffer        ImVector_ImDrawCmd
		IdxBuffer        ImVector_ImDrawIdx
		VtxBuffer        ImVector_ImDrawVert
		Flags            int32
		X_VtxCurrentIdx  uint32
		X_Data           *ImDrawListSharedData
		X_OwnerName      *int8
		X_VtxWritePtr    *ImDrawVert
		X_IdxWritePtr    *uint16
		X_ClipRectStack  ImVector_ImVec4
		X_TextureIdStack ImVector_ImTextureID
		X_Path           ImVector_ImVec2
		X_CmdHeader      ImDrawCmdHeader
		X_Splitter       ImDrawListSplitter
		X_FringeScale    float32
		Pad_cgo_0        [4]byte
	}
	ImGuiPlatformImeData struct {
		WantVisible     bool
		InputPos        ImVec2
		InputLineHeight float32
	}
	ImGuiTableSettings struct {
		ID              uint32
		SaveFlags       int32
		RefScale        float32
		ColumnsCount    int16
		ColumnsCountMax int16
		WantApply       bool
		Pad_cgo_0       [3]byte
	}
	ImGuiTableTempData struct {
		TableIndex                   int32
		LastTimeActive               float32
		UserOuterSize                ImVec2
		DrawSplitter                 ImDrawListSplitter
		HostBackupWorkRect           ImRect
		HostBackupParentWorkRect     ImRect
		HostBackupPrevLineSize       ImVec2
		HostBackupCurrLineSize       ImVec2
		HostBackupCursorMaxPos       ImVec2
		HostBackupColumnsOffset      ImVec1
		HostBackupItemWidth          float32
		HostBackupItemWidthStackSize int32
		Pad_cgo_0                    [4]byte
	}
	ImGuiTableCellData struct {
		BgColor   uint32
		Column    int16
		Pad_cgo_0 [2]byte
	}
	ImFontAtlas struct {
		Flags              int32
		TexID              *byte
		TexDesiredWidth    int32
		TexGlyphPadding    int32
		Locked             bool
		UserData           *byte
		TexReady           bool
		TexPixelsUseColors bool
		TexPixelsAlpha8    *uint8
		TexPixelsRGBA32    *uint32
		TexWidth           int32
		TexHeight          int32
		TexUvScale         ImVec2
		TexUvWhitePixel    ImVec2
		Fonts              ImVector_ImFontPtr
		CustomRects        ImVector_ImFontAtlasCustomRect
		ConfigData         ImVector_ImFontConfig
		TexUvLines         [64]ImVec4
		FontBuilderIO      *ImFontBuilderIO
		FontBuilderFlags   uint32
		PackIdMouseCursors int32
		PackIdLines        int32
		Pad_cgo_0          [4]byte
	}
	ImFontGlyph struct {
		Pad_cgo_0 [4]byte
		AdvanceX  float32
		X0        float32
		Y0        float32
		X1        float32
		Y1        float32
		U0        float32
		V0        float32
		U1        float32
		V1        float32
	}
	ImGuiStyle struct {
		Alpha                      float32
		DisabledAlpha              float32
		WindowPadding              ImVec2
		WindowRounding             float32
		WindowBorderSize           float32
		WindowMinSize              ImVec2
		WindowTitleAlign           ImVec2
		WindowMenuButtonPosition   int32
		ChildRounding              float32
		ChildBorderSize            float32
		PopupRounding              float32
		PopupBorderSize            float32
		FramePadding               ImVec2
		FrameRounding              float32
		FrameBorderSize            float32
		ItemSpacing                ImVec2
		ItemInnerSpacing           ImVec2
		CellPadding                ImVec2
		TouchExtraPadding          ImVec2
		IndentSpacing              float32
		ColumnsMinSpacing          float32
		ScrollbarSize              float32
		ScrollbarRounding          float32
		GrabMinSize                float32
		GrabRounding               float32
		LogSliderDeadzone          float32
		TabRounding                float32
		TabBorderSize              float32
		TabMinWidthForCloseButton  float32
		ColorButtonPosition        int32
		ButtonTextAlign            ImVec2
		SelectableTextAlign        ImVec2
		SeparatorTextBorderSize    float32
		SeparatorTextAlign         ImVec2
		SeparatorTextPadding       ImVec2
		DisplayWindowPadding       ImVec2
		DisplaySafeAreaPadding     ImVec2
		MouseCursorScale           float32
		AntiAliasedLines           bool
		AntiAliasedLinesUseTex     bool
		AntiAliasedFill            bool
		CurveTessellationTol       float32
		CircleTessellationMaxError float32
		Colors                     [55]ImVec4
		HoverStationaryDelay       float32
		HoverDelayShort            float32
		HoverDelayNormal           float32
		HoverFlagsForTooltipMouse  int32
		HoverFlagsForTooltipNav    int32
	}
	ImGuiWindowTempData struct {
		CursorPos                 ImVec2
		CursorPosPrevLine         ImVec2
		CursorStartPos            ImVec2
		CursorMaxPos              ImVec2
		IdealMaxPos               ImVec2
		CurrLineSize              ImVec2
		PrevLineSize              ImVec2
		CurrLineTextBaseOffset    float32
		PrevLineTextBaseOffset    float32
		IsSameLine                bool
		IsSetPos                  bool
		Indent                    ImVec1
		ColumnsOffset             ImVec1
		GroupOffset               ImVec1
		CursorStartPosLossyness   ImVec2
		NavLayerCurrent           uint32
		NavLayersActiveMask       int16
		NavLayersActiveMaskNext   int16
		NavIsScrollPushableX      bool
		NavHideHighlightOneFrame  bool
		NavWindowHasScrollY       bool
		MenuBarAppending          bool
		MenuBarOffset             ImVec2
		MenuColumns               ImGuiMenuColumns
		TreeDepth                 int32
		TreeJumpToParentOnPopMask uint32
		ChildWindows              ImVector_ImGuiWindowPtr
		StateStorage              *ImGuiStorage
		CurrentColumns            *ImGuiOldColumns
		CurrentTableIdx           int32
		LayoutType                int32
		ParentLayoutType          int32
		ItemWidth                 float32
		TextWrapPos               float32
		ItemWidthStack            ImVector_float
		TextWrapPosStack          ImVector_float
	}
	ImGuiInputEvent struct {
		Type              uint32
		Source            uint32
		EventId           uint32
		MousePos          ImGuiInputEventMousePos
		AddedByTestEngine bool
		Pad_cgo_0         [3]byte
	}
	ImVec2 struct {
		X float32
		Y float32
	}
	ImGuiTabItem struct {
		ID                uint32
		Flags             int32
		Window            *ImGuiWindow
		LastFrameVisible  int32
		LastFrameSelected int32
		Offset            float32
		Width             float32
		ContentWidth      float32
		RequestedWidth    float32
		NameOffset        int32
		BeginOrder        int16
		IndexDuringLayout int16
		WantClose         bool
		Pad_cgo_0         [7]byte
	}
	ImGuiTable struct {
		ID                         uint32
		Flags                      int32
		RawData                    *byte
		TempData                   *ImGuiTableTempData
		Columns                    ImSpan_ImGuiTableColumn
		DisplayOrderToIndex        ImSpan_ImGuiTableColumnIdx
		RowCellData                ImSpan_ImGuiTableCellData
		EnabledMaskByDisplayOrder  *uint32
		EnabledMaskByIndex         *uint32
		VisibleMaskByIndex         *uint32
		SettingsLoadedFlags        int32
		SettingsOffset             int32
		LastFrameActive            int32
		ColumnsCount               int32
		CurrentRow                 int32
		CurrentColumn              int32
		InstanceCurrent            int16
		InstanceInteracted         int16
		RowPosY1                   float32
		RowPosY2                   float32
		RowMinHeight               float32
		RowTextBaseline            float32
		RowIndentOffsetX           float32
		Pad_cgo_0                  [4]byte
		RowBgColorCounter          int32
		RowBgColor                 [2]uint32
		BorderColorStrong          uint32
		BorderColorLight           uint32
		BorderX1                   float32
		BorderX2                   float32
		HostIndentX                float32
		MinColumnWidth             float32
		OuterPaddingX              float32
		CellPaddingX               float32
		CellPaddingY               float32
		CellSpacingX1              float32
		CellSpacingX2              float32
		InnerWidth                 float32
		ColumnsGivenWidth          float32
		ColumnsAutoFitWidth        float32
		ColumnsStretchSumWeights   float32
		ResizedColumnNextWidth     float32
		ResizeLockMinContentsX2    float32
		RefScale                   float32
		OuterRect                  ImRect
		InnerRect                  ImRect
		WorkRect                   ImRect
		InnerClipRect              ImRect
		BgClipRect                 ImRect
		Bg0ClipRectForDrawCmd      ImRect
		Bg2ClipRectForDrawCmd      ImRect
		HostClipRect               ImRect
		HostBackupInnerClipRect    ImRect
		OuterWindow                *ImGuiWindow
		InnerWindow                *ImGuiWindow
		ColumnsNames               ImGuiTextBuffer
		DrawSplitter               *ImDrawListSplitter
		InstanceDataFirst          ImGuiTableInstanceData
		InstanceDataExtra          ImVector_ImGuiTableInstanceData
		SortSpecsSingle            ImGuiTableColumnSortSpecs
		SortSpecsMulti             ImVector_ImGuiTableColumnSortSpecs
		SortSpecs                  ImGuiTableSortSpecs
		SortSpecsCount             int16
		ColumnsEnabledCount        int16
		ColumnsEnabledFixedCount   int16
		DeclColumnsCount           int16
		HoveredColumnBody          int16
		HoveredColumnBorder        int16
		AutoFitSingleColumn        int16
		ResizedColumn              int16
		LastResizedColumn          int16
		HeldHeaderColumn           int16
		ReorderColumn              int16
		ReorderColumnDir           int16
		LeftMostEnabledColumn      int16
		RightMostEnabledColumn     int16
		LeftMostStretchedColumn    int16
		RightMostStretchedColumn   int16
		ContextPopupColumn         int16
		FreezeRowsRequest          int16
		FreezeRowsCount            int16
		FreezeColumnsRequest       int16
		FreezeColumnsCount         int16
		RowCellDataCurrent         int16
		DummyDrawChannel           uint16
		Bg2DrawChannelCurrent      uint16
		Bg2DrawChannelUnfrozen     uint16
		IsLayoutLocked             bool
		IsInsideRow                bool
		IsInitializing             bool
		IsSortSpecsDirty           bool
		IsUsingHeaders             bool
		IsContextPopupOpen         bool
		IsSettingsRequestLoad      bool
		IsSettingsDirty            bool
		IsDefaultDisplayOrder      bool
		IsResetAllRequest          bool
		IsResetDisplayOrderRequest bool
		IsUnfrozenRows             bool
		IsDefaultSizingPolicy      bool
		HasScrollbarYCurr          bool
		HasScrollbarYPrev          bool
		MemoryCompacted            bool
		HostSkipItems              bool
		Pad_cgo_1                  [5]byte
	}
	ImGuiTableColumn struct {
		Flags                    int32
		WidthGiven               float32
		MinX                     float32
		MaxX                     float32
		WidthRequest             float32
		WidthAuto                float32
		StretchWeight            float32
		InitStretchWeightOrWidth float32
		ClipRect                 ImRect
		UserID                   uint32
		WorkMinX                 float32
		WorkMaxX                 float32
		ItemWidth                float32
		ContentMaxXFrozen        float32
		ContentMaxXUnfrozen      float32
		ContentMaxXHeadersUsed   float32
		ContentMaxXHeadersIdeal  float32
		NameOffset               int16
		DisplayOrder             int16
		IndexWithinEnabledSet    int16
		PrevEnabledColumn        int16
		NextEnabledColumn        int16
		SortOrder                int16
		DrawChannelCurrent       uint16
		DrawChannelFrozen        uint16
		DrawChannelUnfrozen      uint16
		IsEnabled                bool
		IsUserEnabled            bool
		IsUserEnabledNextFrame   bool
		IsVisibleX               bool
		IsVisibleY               bool
		IsRequestOutput          bool
		IsSkipItems              bool
		IsPreserveWidthAuto      bool
		NavLayerCurrent          int8
		AutoFitQueue             uint8
		CannotSkipItemsQueue     uint8
		Pad_cgo_0                [1]byte
		SortDirectionsAvailList  uint8
		Pad_cgo_1                [1]byte
	}
	ImGuiDockContext struct {
		Nodes           ImGuiStorage
		Requests        ImVector_ImGuiDockRequest
		NodesSettings   ImVector_ImGuiDockNodeSettings
		WantFullRebuild bool
		Pad_cgo_0       [7]byte
	}

	// ImGuiDockRequest _cgopackage.Incomplete
	ImGuiInputTextDeactivatedState struct {
		ID    uint32
		TextA ImVector_char
	}
	ImGuiOldColumnData struct {
		OffsetNorm             float32
		OffsetNormBeforeResize float32
		Flags                  int32
		ClipRect               ImRect
	}
	ImGuiStackTool struct {
		LastActiveFrame         int32
		StackLevel              int32
		QueryId                 uint32
		Results                 ImVector_ImGuiStackLevelInfo
		CopyToClipboardOnCtrlC  bool
		CopyToClipboardLastTime float32
	}
	ImBitArray_ImGuiKey_NamedKey_COUNT__lessImGuiKey_NamedKey_BEGIN struct {
		Storage [5]uint32
	}
	ImGuiMenuColumns struct {
		TotalWidth     uint32
		NextTotalWidth uint32
		Spacing        uint16
		OffsetIcon     uint16
		OffsetLabel    uint16
		OffsetShortcut uint16
		OffsetMark     uint16
		Widths         [4]uint16
		Pad_cgo_0      [2]byte
	}
	ImGuiNavItemData struct {
		Window       *ImGuiWindow
		ID           uint32
		FocusScopeId uint32
		RectRel      ImRect
		InFlags      int32
		DistBox      float32
		DistCenter   float32
		DistAxial    float32
	}
	ImGuiViewport struct {
		ID                    uint32
		Flags                 int32
		Pos                   ImVec2
		Size                  ImVec2
		WorkPos               ImVec2
		WorkSize              ImVec2
		DpiScale              float32
		ParentViewportId      uint32
		DrawData              *ImDrawData
		RendererUserData      *byte
		PlatformUserData      *byte
		PlatformHandle        *byte
		PlatformHandleRaw     *byte
		PlatformWindowCreated bool
		PlatformRequestMove   bool
		PlatformRequestResize bool
		PlatformRequestClose  bool
		Pad_cgo_0             [4]byte
	}
	ImGuiInputTextState struct {
		Ctx                  *ImGuiContext
		ID                   uint32
		CurLenW              int32
		CurLenA              int32
		TextW                ImVector_ImWchar
		TextA                ImVector_char
		InitialTextA         ImVector_char
		TextAIsValid         bool
		BufCapacityA         int32
		ScrollX              float32
		Stb                  STB_TexteditState
		CursorAnim           float32
		CursorFollow         bool
		SelectedAllMouseLock bool
		Edited               bool
		Flags                int32
		Pad_cgo_0            [4]byte
	}
	ImGuiKeyRoutingData struct {
		NextEntryIndex   int16
		Mods             uint16
		RoutingNextScore uint8
		RoutingCurr      uint32
		RoutingNext      uint32
	}
	ImGuiComboPreviewData struct {
		PreviewRect                  ImRect
		BackupCursorPos              ImVec2
		BackupCursorMaxPos           ImVec2
		BackupCursorPosPrevLine      ImVec2
		BackupPrevLineTextBaseOffset float32
		BackupLayout                 int32
	}
	ImGuiContext struct {
		Initialized                              bool
		FontAtlasOwnedByContext                  bool
		IO                                       ImGuiIO
		PlatformIO                               ImGuiPlatformIO
		Style                                    ImGuiStyle
		ConfigFlagsCurrFrame                     int32
		ConfigFlagsLastFrame                     int32
		Font                                     *ImFont
		FontSize                                 float32
		FontBaseSize                             float32
		DrawListSharedData                       ImDrawListSharedData
		Time                                     float64
		FrameCount                               int32
		FrameCountEnded                          int32
		FrameCountPlatformEnded                  int32
		FrameCountRendered                       int32
		WithinFrameScope                         bool
		WithinFrameScopeWithImplicitWindow       bool
		WithinEndChild                           bool
		GcCompactAll                             bool
		TestEngineHookItems                      bool
		TestEngine                               *byte
		InputEventsQueue                         ImVector_ImGuiInputEvent
		InputEventsTrail                         ImVector_ImGuiInputEvent
		InputEventsNextMouseSource               uint32
		InputEventsNextEventId                   uint32
		Windows                                  ImVector_ImGuiWindowPtr
		WindowsFocusOrder                        ImVector_ImGuiWindowPtr
		WindowsTempSortBuffer                    ImVector_ImGuiWindowPtr
		CurrentWindowStack                       ImVector_ImGuiWindowStackData
		WindowsById                              ImGuiStorage
		WindowsActiveCount                       int32
		WindowsHoverPadding                      ImVec2
		CurrentWindow                            *ImGuiWindow
		HoveredWindow                            *ImGuiWindow
		HoveredWindowUnderMovingWindow           *ImGuiWindow
		MovingWindow                             *ImGuiWindow
		WheelingWindow                           *ImGuiWindow
		WheelingWindowRefMousePos                ImVec2
		WheelingWindowStartFrame                 int32
		WheelingWindowReleaseTimer               float32
		WheelingWindowWheelRemainder             ImVec2
		WheelingAxisAvg                          ImVec2
		DebugHookIdInfo                          uint32
		HoveredId                                uint32
		HoveredIdPreviousFrame                   uint32
		HoveredIdAllowOverlap                    bool
		HoveredIdDisabled                        bool
		HoveredIdTimer                           float32
		HoveredIdNotActiveTimer                  float32
		ActiveId                                 uint32
		ActiveIdIsAlive                          uint32
		ActiveIdTimer                            float32
		ActiveIdIsJustActivated                  bool
		ActiveIdAllowOverlap                     bool
		ActiveIdNoClearOnFocusLoss               bool
		ActiveIdHasBeenPressedBefore             bool
		ActiveIdHasBeenEditedBefore              bool
		ActiveIdHasBeenEditedThisFrame           bool
		ActiveIdClickOffset                      ImVec2
		ActiveIdWindow                           *ImGuiWindow
		ActiveIdSource                           uint32
		ActiveIdMouseButton                      int32
		ActiveIdPreviousFrame                    uint32
		ActiveIdPreviousFrameIsAlive             bool
		ActiveIdPreviousFrameHasBeenEditedBefore bool
		ActiveIdPreviousFrameWindow              *ImGuiWindow
		LastActiveId                             uint32
		LastActiveIdTimer                        float32
		KeysOwnerData                            [140]ImGuiKeyOwnerData
		KeysRoutingTable                         ImGuiKeyRoutingTable
		ActiveIdUsingNavDirMask                  uint32
		ActiveIdUsingAllKeyboardKeys             bool
		ActiveIdUsingNavInputMask                uint32
		CurrentFocusScopeId                      uint32
		CurrentItemFlags                         int32
		DebugLocateId                            uint32
		NextItemData                             ImGuiNextItemData
		LastItemData                             ImGuiLastItemData
		NextWindowData                           ImGuiNextWindowData
		ColorStack                               ImVector_ImGuiColorMod
		StyleVarStack                            ImVector_ImGuiStyleMod
		FontStack                                ImVector_ImFontPtr
		FocusScopeStack                          ImVector_ImGuiID
		ItemFlagsStack                           ImVector_ImGuiItemFlags
		GroupStack                               ImVector_ImGuiGroupData
		OpenPopupStack                           ImVector_ImGuiPopupData
		BeginPopupStack                          ImVector_ImGuiPopupData
		BeginMenuCount                           int32
		Viewports                                ImVector_ImGuiViewportPPtr
		CurrentDpiScale                          float32
		CurrentViewport                          *ImGuiViewportP
		MouseViewport                            *ImGuiViewportP
		MouseLastHoveredViewport                 *ImGuiViewportP
		PlatformLastFocusedViewportId            uint32
		FallbackMonitor                          ImGuiPlatformMonitor
		ViewportCreatedCount                     int32
		PlatformWindowsCreatedCount              int32
		ViewportFocusedStampCount                int32
		NavWindow                                *ImGuiWindow
		NavId                                    uint32
		NavFocusScopeId                          uint32
		NavActivateId                            uint32
		NavActivateDownId                        uint32
		NavActivatePressedId                     uint32
		NavActivateFlags                         int32
		NavJustMovedToId                         uint32
		NavJustMovedToFocusScopeId               uint32
		NavJustMovedToKeyMods                    int32
		NavNextActivateId                        uint32
		NavNextActivateFlags                     int32
		NavInputSource                           uint32
		NavLayer                                 uint32
		NavIdIsAlive                             bool
		NavMousePosDirty                         bool
		NavDisableHighlight                      bool
		NavDisableMouseHover                     bool
		NavAnyRequest                            bool
		NavInitRequest                           bool
		NavInitRequestFromMove                   bool
		NavInitResult                            ImGuiNavItemData
		NavMoveSubmitted                         bool
		NavMoveScoringItems                      bool
		NavMoveForwardToNextFrame                bool
		NavMoveFlags                             int32
		NavMoveScrollFlags                       int32
		NavMoveKeyMods                           int32
		NavMoveDir                               int32
		NavMoveDirForDebug                       int32
		NavMoveClipDir                           int32
		NavScoringRect                           ImRect
		NavScoringNoClipRect                     ImRect
		NavScoringDebugCount                     int32
		NavTabbingDir                            int32
		NavTabbingCounter                        int32
		NavMoveResultLocal                       ImGuiNavItemData
		NavMoveResultLocalVisible                ImGuiNavItemData
		NavMoveResultOther                       ImGuiNavItemData
		NavTabbingResultFirst                    ImGuiNavItemData
		ConfigNavWindowingKeyNext                int32
		ConfigNavWindowingKeyPrev                int32
		NavWindowingTarget                       *ImGuiWindow
		NavWindowingTargetAnim                   *ImGuiWindow
		NavWindowingListWindow                   *ImGuiWindow
		NavWindowingTimer                        float32
		NavWindowingHighlightAlpha               float32
		NavWindowingToggleLayer                  bool
		NavWindowingAccumDeltaPos                ImVec2
		NavWindowingAccumDeltaSize               ImVec2
		DimBgRatio                               float32
		DragDropActive                           bool
		DragDropWithinSource                     bool
		DragDropWithinTarget                     bool
		DragDropSourceFlags                      int32
		DragDropSourceFrameCount                 int32
		DragDropMouseButton                      int32
		DragDropPayload                          ImGuiPayload
		DragDropTargetRect                       ImRect
		DragDropTargetId                         uint32
		DragDropAcceptFlags                      int32
		DragDropAcceptIdCurrRectSurface          float32
		DragDropAcceptIdCurr                     uint32
		DragDropAcceptIdPrev                     uint32
		DragDropAcceptFrameCount                 int32
		DragDropHoldJustPressedId                uint32
		DragDropPayloadBufHeap                   ImVector_unsigned_char
		DragDropPayloadBufLocal                  [16]uint8
		ClipperTempDataStacked                   int32
		ClipperTempData                          ImVector_ImGuiListClipperData
		CurrentTable                             *ImGuiTable
		TablesTempDataStacked                    int32
		TablesTempData                           ImVector_ImGuiTableTempData
		Tables                                   ImPool_ImGuiTable
		TablesLastTimeActive                     ImVector_float
		DrawChannelsTempMergeBuffer              ImVector_ImDrawChannel
		CurrentTabBar                            *ImGuiTabBar
		TabBars                                  ImPool_ImGuiTabBar
		CurrentTabBarStack                       ImVector_ImGuiPtrOrIndex
		ShrinkWidthBuffer                        ImVector_ImGuiShrinkWidthItem
		HoverItemDelayId                         uint32
		HoverItemDelayIdPreviousFrame            uint32
		HoverItemDelayTimer                      float32
		HoverItemDelayClearTimer                 float32
		HoverItemUnlockedStationaryId            uint32
		HoverWindowUnlockedStationaryId          uint32
		MouseCursor                              int32
		MouseStationaryTimer                     float32
		MouseLastValidPos                        ImVec2
		InputTextState                           ImGuiInputTextState
		InputTextDeactivatedState                ImGuiInputTextDeactivatedState
		InputTextPasswordFont                    ImFont
		TempInputId                              uint32
		ColorEditOptions                         int32
		ColorEditCurrentID                       uint32
		ColorEditSavedID                         uint32
		ColorEditSavedHue                        float32
		ColorEditSavedSat                        float32
		ColorEditSavedColor                      uint32
		ColorPickerRef                           ImVec4
		ComboPreviewData                         ImGuiComboPreviewData
		SliderGrabClickOffset                    float32
		SliderCurrentAccum                       float32
		SliderCurrentAccumDirty                  bool
		DragCurrentAccumDirty                    bool
		DragCurrentAccum                         float32
		DragSpeedDefaultRatio                    float32
		ScrollbarClickDeltaToGrabCenter          float32
		DisabledAlphaBackup                      float32
		DisabledStackSize                        int16
		TooltipOverrideCount                     int16
		ClipboardHandlerData                     ImVector_char
		MenusIdSubmittedThisFrame                ImVector_ImGuiID
		PlatformImeData                          ImGuiPlatformImeData
		PlatformImeDataPrev                      ImGuiPlatformImeData
		PlatformImeViewport                      uint32
		PlatformLocaleDecimalPoint               int8
		DockContext                              ImGuiDockContext
		DockNodeWindowMenuHandler                *[0]byte
		SettingsLoaded                           bool
		SettingsDirtyTimer                       float32
		SettingsIniData                          ImGuiTextBuffer
		SettingsHandlers                         ImVector_ImGuiSettingsHandler
		SettingsWindows                          ImChunkStream_ImGuiWindowSettings
		SettingsTables                           ImChunkStream_ImGuiTableSettings
		Hooks                                    ImVector_ImGuiContextHook
		HookIdNext                               uint32
		LocalizationTable                        [9]*int8
		LogEnabled                               bool
		LogType                                  uint32
		LogFile                                  ImFileHandle
		LogBuffer                                ImGuiTextBuffer
		LogNextPrefix                            *int8
		LogNextSuffix                            *int8
		LogLinePosY                              float32
		LogLineFirstItem                         bool
		LogDepthRef                              int32
		LogDepthToExpand                         int32
		LogDepthToExpandDefault                  int32
		DebugLogFlags                            int32
		DebugLogBuf                              ImGuiTextBuffer
		DebugLogIndex                            ImGuiTextIndex
		DebugLogClipperAutoDisableFrames         uint8
		DebugLocateFrames                        uint8
		DebugBeginReturnValueCullDepth           int8
		DebugItemPickerActive                    bool
		DebugItemPickerMouseButton               uint8
		DebugItemPickerBreakId                   uint32
		DebugMetricsConfig                       ImGuiMetricsConfig
		DebugStackTool                           ImGuiStackTool
		DebugHoveredDockNode                     *ImGuiDockNode
		FramerateSecPerFrame                     [60]float32
		FramerateSecPerFrameIdx                  int32
		FramerateSecPerFrameCount                int32
		FramerateSecPerFrameAccum                float32
		WantCaptureMouseNextFrame                int32
		WantCaptureKeyboardNextFrame             int32
		WantTextInputNextFrame                   int32
		TempBuffer                               ImVector_char
	}
	ImGuiTableColumnSettings struct {
		WidthOrWeight float32
		UserID        uint32
		Index         int16
		DisplayOrder  int16
		SortOrder     int16
		Pad_cgo_0     [2]byte
	}
	ImGuiWindow struct {
		Ctx                                *ImGuiContext
		Name                               *int8
		ID                                 uint32
		Flags                              int32
		FlagsPreviousFrame                 int32
		WindowClass                        ImGuiWindowClass
		Viewport                           *ImGuiViewportP
		ViewportId                         uint32
		ViewportPos                        ImVec2
		ViewportAllowPlatformMonitorExtend int32
		Pos                                ImVec2
		Size                               ImVec2
		SizeFull                           ImVec2
		ContentSize                        ImVec2
		ContentSizeIdeal                   ImVec2
		ContentSizeExplicit                ImVec2
		WindowPadding                      ImVec2
		WindowRounding                     float32
		WindowBorderSize                   float32
		DecoOuterSizeX1                    float32
		DecoOuterSizeY1                    float32
		DecoOuterSizeX2                    float32
		DecoOuterSizeY2                    float32
		DecoInnerSizeX1                    float32
		DecoInnerSizeY1                    float32
		NameBufLen                         int32
		MoveId                             uint32
		TabId                              uint32
		ChildId                            uint32
		Scroll                             ImVec2
		ScrollMax                          ImVec2
		ScrollTarget                       ImVec2
		ScrollTargetCenterRatio            ImVec2
		ScrollTargetEdgeSnapDist           ImVec2
		ScrollbarSizes                     ImVec2
		ScrollbarX                         bool
		ScrollbarY                         bool
		ViewportOwned                      bool
		Active                             bool
		WasActive                          bool
		WriteAccessed                      bool
		Collapsed                          bool
		WantCollapseToggle                 bool
		SkipItems                          bool
		Appearing                          bool
		Hidden                             bool
		IsFallbackWindow                   bool
		IsExplicitChild                    bool
		HasCloseButton                     bool
		ResizeBorderHeld                   int8
		BeginCount                         int16
		BeginCountPreviousFrame            int16
		BeginOrderWithinParent             int16
		BeginOrderWithinContext            int16
		FocusOrder                         int16
		PopupId                            uint32
		AutoFitFramesX                     int8
		AutoFitFramesY                     int8
		AutoFitChildAxises                 int8
		AutoFitOnlyGrows                   bool
		AutoPosLastDirection               int32
		HiddenFramesCanSkipItems           int8
		HiddenFramesCannotSkipItems        int8
		HiddenFramesForRenderOnly          int8
		DisableInputsFrames                int8
		Pad_cgo_0                          [4]byte
		SetWindowPosVal                    ImVec2
		SetWindowPosPivot                  ImVec2
		IDStack                            ImVector_ImGuiID
		DC                                 ImGuiWindowTempData
		OuterRectClipped                   ImRect
		InnerRect                          ImRect
		InnerClipRect                      ImRect
		WorkRect                           ImRect
		ParentWorkRect                     ImRect
		ClipRect                           ImRect
		ContentRegionRect                  ImRect
		HitTestHoleSize                    ImVec2ih
		HitTestHoleOffset                  ImVec2ih
		LastFrameActive                    int32
		LastFrameJustFocused               int32
		LastTimeActive                     float32
		ItemWidthDefault                   float32
		StateStorage                       ImGuiStorage
		ColumnsStorage                     ImVector_ImGuiOldColumns
		FontWindowScale                    float32
		FontDpiScale                       float32
		SettingsOffset                     int32
		DrawList                           *ImDrawList
		DrawListInst                       ImDrawList
		ParentWindow                       *ImGuiWindow
		ParentWindowInBeginStack           *ImGuiWindow
		RootWindow                         *ImGuiWindow
		RootWindowPopupTree                *ImGuiWindow
		RootWindowDockTree                 *ImGuiWindow
		RootWindowForTitleBarHighlight     *ImGuiWindow
		RootWindowForNav                   *ImGuiWindow
		NavLastChildNavWindow              *ImGuiWindow
		NavLastIds                         [2]uint32
		NavRectRel                         [2]ImRect
		NavPreferredScoringPosRel          [2]ImVec2
		NavRootFocusScopeId                uint32
		MemoryDrawListIdxCapacity          int32
		MemoryDrawListVtxCapacity          int32
		MemoryCompacted                    bool
		DockOrder                          int16
		DockStyle                          ImGuiWindowDockStyle
		DockNode                           *ImGuiDockNode
		DockNodeAsHost                     *ImGuiDockNode
		DockId                             uint32
		DockTabItemStatusFlags             int32
		DockTabItemRect                    ImRect
	}
	ImGuiColorMod struct {
		Col         int32
		BackupValue ImVec4
	}
	ImGuiWindowClass struct {
		ClassId                    uint32
		ParentViewportId           uint32
		ViewportFlagsOverrideSet   int32
		ViewportFlagsOverrideClear int32
		TabItemFlagsOverrideSet    int32
		DockNodeFlagsOverrideSet   int32
		DockingAlwaysTabBar        bool
		DockingAllowUnclassed      bool
		Pad_cgo_0                  [2]byte
	}
	ImGuiDataTypeInfo struct {
		Size     uint64
		Name     *int8
		PrintFmt *int8
		ScanFmt  *int8
	}
	ImGuiDockNode struct {
		ID                   uint32
		SharedFlags          int32
		LocalFlags           int32
		LocalFlagsInWindows  int32
		MergedFlags          int32
		State                uint32
		ParentNode           *ImGuiDockNode
		ChildNodes           [2]*ImGuiDockNode
		Windows              ImVector_ImGuiWindowPtr
		TabBar               *ImGuiTabBar
		Pos                  ImVec2
		Size                 ImVec2
		SizeRef              ImVec2
		SplitAxis            int32
		WindowClass          ImGuiWindowClass
		LastBgColor          uint32
		HostWindow           *ImGuiWindow
		VisibleWindow        *ImGuiWindow
		CentralNode          *ImGuiDockNode
		OnlyNodeWithWindows  *ImGuiDockNode
		CountNodeWithWindows int32
		LastFrameAlive       int32
		LastFrameActive      int32
		LastFrameFocused     int32
		LastFocusedNodeId    uint32
		SelectedTabId        uint32
		WantCloseTabId       uint32
		RefViewportId        uint32
		Pad_cgo_0            [8]byte
	}
	ImGuiPlatformIO struct {
		Platform_CreateWindow       *[0]byte
		Platform_DestroyWindow      *[0]byte
		Platform_ShowWindow         *[0]byte
		Platform_SetWindowPos       *[0]byte
		Platform_GetWindowPos       *[0]byte
		Platform_SetWindowSize      *[0]byte
		Platform_GetWindowSize      *[0]byte
		Platform_SetWindowFocus     *[0]byte
		Platform_GetWindowFocus     *[0]byte
		Platform_GetWindowMinimized *[0]byte
		Platform_SetWindowTitle     *[0]byte
		Platform_SetWindowAlpha     *[0]byte
		Platform_UpdateWindow       *[0]byte
		Platform_RenderWindow       *[0]byte
		Platform_SwapBuffers        *[0]byte
		Platform_GetWindowDpiScale  *[0]byte
		Platform_OnChangedViewport  *[0]byte
		Platform_CreateVkSurface    *[0]byte
		Renderer_CreateWindow       *[0]byte
		Renderer_DestroyWindow      *[0]byte
		Renderer_SetWindowSize      *[0]byte
		Renderer_RenderWindow       *[0]byte
		Renderer_SwapBuffers        *[0]byte
		Monitors                    ImVector_ImGuiPlatformMonitor
		Viewports                   ImVector_ImGuiViewportPtr
	}
	ImGuiShrinkWidthItem struct {
		Index        int32
		Width        float32
		InitialWidth float32
	}
	ImGuiStorage struct {
		Data ImVector_ImGuiStoragePair
	}
	ImDrawListSharedData struct {
		TexUvWhitePixel       ImVec2
		Font                  *ImFont
		FontSize              float32
		CurveTessellationTol  float32
		CircleSegmentMaxError float32
		ClipRectFullscreen    ImVec4
		InitialFlags          int32
		TempBuffer            ImVector_ImVec2
		ArcFastVtx            [48]ImVec2
		ArcFastRadiusCutoff   float32
		CircleSegmentCounts   [64]uint8
		TexUvLines            *ImVec4
	}
	ImGuiInputEventMouseViewport struct {
		HoveredViewportID uint32
	}
	ImGuiLocEntry struct {
		Key  uint32
		Text *int8
	}
	ImGuiListClipper struct {
		Ctx          *ImGuiContext
		DisplayStart int32
		DisplayEnd   int32
		ItemsCount   int32
		ItemsHeight  float32
		StartPosY    float32
		TempData     *byte
	}
	ImColor struct {
		Value ImVec4
	}
	ImGuiStackLevelInfo struct {
		ID              uint32
		QueryFrameCount int8
		QuerySuccess    bool
		Pad_cgo_0       [6]byte
		Desc            [57]int8
		Pad_cgo_1       [3]byte
	}
	ImGuiPlatformMonitor struct {
		MainPos        ImVec2
		MainSize       ImVec2
		WorkPos        ImVec2
		WorkSize       ImVec2
		DpiScale       float32
		PlatformHandle *byte
	}
	ImGuiTabBar struct {
		Tabs                            ImVector_ImGuiTabItem
		Flags                           int32
		ID                              uint32
		SelectedTabId                   uint32
		NextSelectedTabId               uint32
		VisibleTabId                    uint32
		CurrFrameVisible                int32
		PrevFrameVisible                int32
		BarRect                         ImRect
		CurrTabsContentsHeight          float32
		PrevTabsContentsHeight          float32
		WidthAllTabs                    float32
		WidthAllTabsIdeal               float32
		ScrollingAnim                   float32
		ScrollingTarget                 float32
		ScrollingTargetDistToVisibility float32
		ScrollingSpeed                  float32
		ScrollingRectMinX               float32
		ScrollingRectMaxX               float32
		ReorderRequestTabId             uint32
		ReorderRequestOffset            int16
		BeginCount                      int8
		WantLayout                      bool
		VisibleTabWasSubmitted          bool
		TabsAddedNew                    bool
		TabsActiveCount                 int16
		LastTabItemIdx                  int16
		ItemSpacingY                    float32
		FramePadding                    ImVec2
		BackupCursorPos                 ImVec2
		TabsNames                       ImGuiTextBuffer
	}
	ImGuiWindowSettings struct {
		ID          uint32
		Pos         ImVec2ih
		Size        ImVec2ih
		ViewportPos ImVec2ih
		ViewportId  uint32
		DockId      uint32
		ClassId     uint32
		DockOrder   int16
		Collapsed   bool
		WantApply   bool
		WantDelete  bool
		Pad_cgo_0   [3]byte
	}
	ImGuiStackSizes struct {
		SizeOfIDStack         int16
		SizeOfColorStack      int16
		SizeOfStyleVarStack   int16
		SizeOfFontStack       int16
		SizeOfFocusScopeStack int16
		SizeOfGroupStack      int16
		SizeOfItemFlagsStack  int16
		SizeOfBeginPopupStack int16
		SizeOfDisabledStack   int16
	}
	ImGuiStyleMod struct {
		VarIdx    int32
		BackupInt [2]int32
	}
	ImRect struct {
		Min ImVec2
		Max ImVec2
	}
	ImGuiInputEventMouseButton struct {
		Button      int32
		Down        bool
		MouseSource uint32
	}
	ImGuiKeyData struct {
		Down             bool
		DownDuration     float32
		DownDurationPrev float32
		AnalogValue      float32
	}

	// ImGuiInputTextDeactivateData _cgopackage.Incomplete
	ImGuiKeyOwnerData struct {
		OwnerCurr        uint32
		OwnerNext        uint32
		LockThisFrame    bool
		LockUntilRelease bool
		Pad_cgo_0        [2]byte
	}
	ImGuiStoragePair struct {
		Key       uint32
		Pad_cgo_0 [4]byte
		I         int32
		Pad_cgo_1 [4]byte
	}
	ImGuiInputEventAppFocused struct {
		Focused bool
	}
	ImGuiPtrOrIndex struct {
		Ptr       *byte
		Index     int32
		Pad_cgo_0 [4]byte
	}
	ImGuiTableSortSpecs struct {
		Specs      *ImGuiTableColumnSortSpecs
		SpecsCount int32
		SpecsDirty bool
		Pad_cgo_0  [3]byte
	}
	ImGuiTextFilter struct {
		InputBuf  [256]int8
		Filters   ImVector_ImGuiTextRange
		CountGrep int32
		Pad_cgo_0 [4]byte
	}
	ImBitVector struct {
		Storage ImVector_ImU32
	}
	ImGuiDataVarInfo struct {
		Type   int32
		Count  uint32
		Offset uint32
	}
	ImGuiLastItemData struct {
		ID          uint32
		InFlags     int32
		StatusFlags int32
		Rect        ImRect
		NavRect     ImRect
		DisplayRect ImRect
	}
	ImGuiNextItemData struct {
		Flags        int32
		ItemFlags    int32
		Width        float32
		FocusScopeId uint32
		OpenCond     int32
		OpenVal      bool
		Pad_cgo_0    [3]byte
	}
	ImGuiViewportP struct {
		X_ImGuiViewport         ImGuiViewport
		Window                  *ImGuiWindow
		Idx                     int32
		LastFrameActive         int32
		LastFocusedStampCount   int32
		LastNameHash            uint32
		LastPos                 ImVec2
		Alpha                   float32
		LastAlpha               float32
		LastFocusedHadNavWindow bool
		PlatformMonitor         int16
		DrawListsLastFrame      [2]int32
		DrawLists               [2]*ImDrawList
		DrawDataP               ImDrawData
		DrawDataBuilder         ImDrawDataBuilder
		LastPlatformPos         ImVec2
		LastPlatformSize        ImVec2
		LastRendererSize        ImVec2
		WorkOffsetMin           ImVec2
		WorkOffsetMax           ImVec2
		BuildWorkOffsetMin      ImVec2
		BuildWorkOffsetMax      ImVec2
	}
	ImGuiGroupData struct {
		WindowID                           uint32
		BackupCursorPos                    ImVec2
		BackupCursorMaxPos                 ImVec2
		BackupIndent                       ImVec1
		BackupGroupOffset                  ImVec1
		BackupCurrLineSize                 ImVec2
		BackupCurrLineTextBaseOffset       float32
		BackupActiveIdIsAlive              uint32
		BackupActiveIdPreviousFrameIsAlive bool
		BackupHoveredIdIsAlive             bool
		EmitItem                           bool
		Pad_cgo_0                          [1]byte
	}
	ImGuiIO struct {
		ConfigFlags                       ImGuiConfigFlags
		BackendFlags                      int32
		DisplaySize                       ImVec2
		DeltaTime                         float32
		IniSavingRate                     float32
		IniFilename                       *int8
		LogFilename                       *int8
		UserData                          *byte
		Fonts                             *ImFontAtlas
		FontGlobalScale                   float32
		FontAllowUserScaling              bool
		FontDefault                       *ImFont
		DisplayFramebufferScale           ImVec2
		ConfigDockingNoSplit              bool
		ConfigDockingWithShift            bool
		ConfigDockingAlwaysTabBar         bool
		ConfigDockingTransparentPayload   bool
		ConfigViewportsNoAutoMerge        bool
		ConfigViewportsNoTaskBarIcon      bool
		ConfigViewportsNoDecoration       bool
		ConfigViewportsNoDefaultParent    bool
		MouseDrawCursor                   bool
		ConfigMacOSXBehaviors             bool
		ConfigInputTrickleEventQueue      bool
		ConfigInputTextCursorBlink        bool
		ConfigInputTextEnterKeepActive    bool
		ConfigDragClickToInputText        bool
		ConfigWindowsResizeFromEdges      bool
		ConfigWindowsMoveFromTitleBarOnly bool
		ConfigMemoryCompactTimer          float32
		MouseDoubleClickTime              float32
		MouseDoubleClickMaxDist           float32
		MouseDragThreshold                float32
		KeyRepeatDelay                    float32
		KeyRepeatRate                     float32
		ConfigDebugBeginReturnValueOnce   bool
		ConfigDebugBeginReturnValueLoop   bool
		ConfigDebugIgnoreFocusLoss        bool
		ConfigDebugIniSettings            bool
		BackendPlatformName               *int8
		BackendRendererName               *int8
		BackendPlatformUserData           *byte
		BackendRendererUserData           *byte
		BackendLanguageUserData           *byte
		GetClipboardTextFn                *[0]byte
		SetClipboardTextFn                *[0]byte
		ClipboardUserData                 *byte
		SetPlatformImeDataFn              *[0]byte
		X_UnusedPadding                   *byte
		WantCaptureMouse                  bool
		WantCaptureKeyboard               bool
		WantTextInput                     bool
		WantSetMousePos                   bool
		WantSaveIniSettings               bool
		NavActive                         bool
		NavVisible                        bool
		Framerate                         float32
		MetricsRenderVertices             int32
		MetricsRenderIndices              int32
		MetricsRenderWindows              int32
		MetricsActiveWindows              int32
		MetricsActiveAllocations          int32
		MouseDelta                        ImVec2
		KeyMap                            [652]int32
		KeysDown                          [652]bool
		NavInputs                         [16]float32
		Ctx                               *ImGuiContext
		MousePos                          ImVec2
		MouseDown                         [5]bool
		MouseWheel                        float32
		MouseWheelH                       float32
		MouseSource                       uint32
		MouseHoveredViewport              uint32
		KeyCtrl                           bool
		KeyShift                          bool
		KeyAlt                            bool
		KeySuper                          bool
		KeyMods                           int32
		KeysData                          [652]ImGuiKeyData
		WantCaptureMouseUnlessPopupClose  bool
		MousePosPrev                      ImVec2
		MouseClickedPos                   [5]ImVec2
		MouseClickedTime                  [5]float64
		MouseClicked                      [5]bool
		MouseDoubleClicked                [5]bool
		MouseClickedCount                 [5]uint16
		MouseClickedLastCount             [5]uint16
		MouseReleased                     [5]bool
		MouseDownOwned                    [5]bool
		MouseDownOwnedUnlessPopupClose    [5]bool
		MouseWheelRequestAxisSwap         bool
		MouseDownDuration                 [5]float32
		MouseDownDurationPrev             [5]float32
		MouseDragMaxDistanceAbs           [5]ImVec2
		MouseDragMaxDistanceSqr           [5]float32
		PenPressure                       float32
		AppFocusLost                      bool
		AppAcceptingEvents                bool
		BackendUsingLegacyKeyArrays       int8
		BackendUsingLegacyNavInputArray   bool
		InputQueueSurrogate               uint16
		InputQueueCharacters              ImVector_ImWchar
	}
	ImGuiTextRange struct {
		B *int8
		E *int8
	}
	ImGuiWindowStackData struct {
		Window                   *ImGuiWindow
		ParentLastItemDataBackup ImGuiLastItemData
		StackSizesOnBegin        ImGuiStackSizes
		Pad_cgo_0                [2]byte
	}
	ImGuiWindowDockStyle struct {
		Colors [6]uint32
	}
	ImDrawCmd struct {
		ClipRect         ImVec4
		TextureId        *byte
		VtxOffset        uint32
		IdxOffset        uint32
		ElemCount        uint32
		UserCallback     *[0]byte
		UserCallbackData *byte
	}
	ImDrawVert struct {
		Pos ImVec2
		Uv  ImVec2
		Col uint32
	}
	ImGuiContextHook struct {
		HookId   uint32
		Type     uint32
		Owner    uint32
		Callback unsafe.Pointer
		UserData *byte
	}
	ImGuiOnceUponAFrame struct {
		RefFrame int32
	}
	ImGuiPopupData struct {
		PopupId         uint32
		Window          *ImGuiWindow
		BackupNavWindow *ImGuiWindow
		ParentNavLayer  int32
		OpenFrameCount  int32
		OpenParentId    uint32
		OpenPopupPos    ImVec2
		OpenMousePos    ImVec2
		Pad_cgo_0       [4]byte
	}
)

//#####################################

var (
	_TypVec2             = c.TypeStruct(usf.SizeOf(ImVec2{}), uint16(unsafe.Alignof(ImVec2{})), []c.Type{c.F32, c.F32})
	_TypVec4             = c.TypeStruct(usf.SizeOf(ImVec4{}), uint16(unsafe.Alignof(ImVec4{})), []c.Type{c.F32, c.F32, c.F32, c.F32})
	_TypRect             = c.TypeStruct(usf.SizeOf(ImRect{}), uint16(unsafe.Alignof(ImRect{})), []c.Type{_TypVec2, _TypVec2})
	_TypListCliooerRange = c.TypeStruct(usf.SizeOf(ImGuiListClipperRange{}), uint16(unsafe.Alignof(ImRect{})), []c.Type{_TypVec2, _TypVec2})

	_typs_F32                                   = []c.Type{c.F32}
	_typs_F32F32                                = []c.Type{c.F32, c.F32}
	_typs_F32F32F32F32                          = []c.Type{c.F32, c.F32, c.F32, c.F32}
	_typs_F32Vec2P                              = []c.Type{c.F32, _TypVec2, c.Pointer}
	_typs_I32                                   = []c.Type{c.I32}
	_typs_I32P                                  = []c.Type{c.I32, c.Pointer}
	_typs_I32PI32                               = []c.Type{c.I32, c.Pointer, c.I32}
	_typs_I32F32                                = []c.Type{c.I32, c.F32}
	_typs_I32U32                                = []c.Type{c.I32, c.U32}
	_typs_I32Vec2                               = []c.Type{c.I32, _TypVec2}
	_typs_P                                     = []c.Type{c.Pointer}
	_typs_PPP                                   = []c.Type{c.Pointer, c.Pointer, c.Pointer}
	_typs_PF64F64                               = []c.Type{c.Pointer, c.F64, c.F64}
	_typs_PI32F32                               = []c.Type{c.Pointer, c.I32, c.F32}
	_typs_PI32                                  = []c.Type{c.Pointer, c.I32}
	_typs_PI32PP                                = []c.Type{c.Pointer, c.I32, c.Pointer, c.Pointer}
	_typs_PU8                                   = []c.Type{c.Pointer, c.U8}
	_typs_PPF32                                 = []c.Type{c.Pointer, c.Pointer, c.F32}
	_typs_PI32U32Vec2                           = []c.Type{c.Pointer, c.I32, c.U32, _TypVec2}
	_typs_PI32U32Vec2F32                        = []c.Type{c.Pointer, c.I32, c.U32, _TypVec2, c.F32}
	_typs_PPU32Vec2                             = []c.Type{c.Pointer, c.Pointer, c.U32, _TypVec2}
	_typs_PI32I32I32                            = []c.Type{c.Pointer, c.I32, c.I32, c.I32}
	_typs_PI32I32                               = []c.Type{c.Pointer, c.I32, c.I32}
	_typs_I32I32                                = []c.Type{c.I32, c.I32}
	_typs_U32U32I32                             = []c.Type{c.U32, c.U32, c.I32}
	_typs_U32I32                                = []c.Type{c.U32, c.I32}
	_typs_PI32I32I32I32                         = []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32}
	_typs_PP                                    = []c.Type{c.Pointer, c.Pointer}
	_typs_PU8P                                  = []c.Type{c.Pointer, c.U8, c.Pointer}
	_typs_PU32P                                 = []c.Type{c.Pointer, c.U32, c.Pointer}
	_typs_PU32F32U32                            = []c.Type{c.Pointer, c.U32, c.F32, c.U32}
	_typs_PPPU32                                = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_PPF32F32F32PU32                       = []c.Type{c.Pointer, c.Pointer, c.F32, c.F32, c.F32, c.Pointer, c.U32}
	_typs_F32F32F32PPP                          = []c.Type{c.F32, c.F32, c.F32, c.Pointer, c.Pointer, c.Pointer}
	_typs_PPF32F32PU32                          = []c.Type{c.Pointer, c.Pointer, c.F32, c.F32, c.Pointer, c.U32}
	_typs_PPF64F64PU32                          = []c.Type{c.Pointer, c.Pointer, c.F64, c.F64, c.Pointer, c.U32}
	_typs_PPI32I32PU32                          = []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.Pointer, c.U32}
	_typs_PPPI32I32PF32F32Vec2I32               = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.I32, c.I32, c.Pointer, c.F32, c.F32, _TypVec2, c.I32}
	_typs_PPPI32I32PF32F32Vec2                  = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.I32, c.I32, c.Pointer, c.F32, c.F32, _TypVec2}
	_typs_PPI32I32PF32F32Vec2I32                = []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.Pointer, c.F32, c.F32, _TypVec2, c.I32}
	_typs_PPI32I32U32                           = []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.U32}
	_typs_PPF32I32I32PU32                       = []c.Type{c.Pointer, c.Pointer, c.F32, c.I32, c.I32, c.Pointer, c.U32}
	_typs_PPF32PP                               = []c.Type{c.Pointer, c.Pointer, c.F32, c.Pointer, c.Pointer}
	_typs_PF32P                                 = []c.Type{c.Pointer, c.F32, c.Pointer}
	_typs_PF32                                  = []c.Type{c.Pointer, c.F32}
	_typs_PF32F32                               = []c.Type{c.Pointer, c.F32, c.F32}
	_typs_PPI32                                 = []c.Type{c.Pointer, c.Pointer, c.I32}
	_typs_PPI32I32                              = []c.Type{c.Pointer, c.Pointer, c.I32, c.I32}
	_typs_PPI32F32PP                            = []c.Type{c.Pointer, c.Pointer, c.I32, c.F32, c.Pointer, c.Pointer}
	_typs_PPPF32F32F32PPU32                     = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.F32, c.F32, c.F32, c.Pointer, c.Pointer, c.U32}
	_typs_PPPF32I32I32PPU32                     = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.F32, c.I32, c.I32, c.Pointer, c.Pointer, c.U32}
	_typs_PPPI32                                = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.I32}
	_typs_PPPI32F32                             = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.I32, c.F32}
	_typs_PPPI32I32                             = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.I32, c.I32}
	_typs_PPPPI32I32                            = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.I32, c.I32}
	_typs_PPU32                                 = []c.Type{c.Pointer, c.Pointer, c.U32}
	_typs_PPU32P                                = []c.Type{c.Pointer, c.Pointer, c.U32, c.Pointer}
	_typs_PVec4U32Vec2                          = []c.Type{c.Pointer, _TypVec4, c.U32, _TypVec2}
	_typs_PPVec2Vec2Vec2Vec4Vec4                = []c.Type{c.Pointer, c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec4, _TypVec4}
	_typs_PU32                                  = []c.Type{c.Pointer, c.U32}
	_typs_PU16                                  = []c.Type{c.Pointer, c.U16}
	_typs_PU64                                  = []c.Type{c.Pointer, c.U64}
	_typs_PU64U64U64U64U64U64                   = []c.Type{c.Pointer, c.U64, c.U64, c.U64, c.U64, c.U64, c.U64}
	_typs_PU32PPPPU32                           = []c.Type{c.Pointer, c.U32, c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_PVec2U32PPPPU32                       = []c.Type{c.Pointer, _TypVec2, c.U32, c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_PPU64U32PP                            = []c.Type{c.Pointer, c.Pointer, c.U64, c.U32, c.Pointer, c.Pointer}
	_typs_PPU64U32                              = []c.Type{c.Pointer, c.Pointer, c.U64, c.U32}
	_typs_PPPU64U32PP                           = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.U64, c.U32, c.Pointer, c.Pointer}
	_typs_PPU64Vec2U32PP                        = []c.Type{c.Pointer, c.Pointer, c.U64, _TypVec2, c.U32, c.Pointer, c.Pointer}
	_typs_PU32PF32PPPU32                        = []c.Type{c.Pointer, c.U32, c.Pointer, c.F32, c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_PU32PI32PPPU32                        = []c.Type{c.Pointer, c.U32, c.Pointer, c.I32, c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_PU32I32                               = []c.Type{c.Pointer, c.U32, c.I32}
	_typs_PU32I32I32I32                         = []c.Type{c.Pointer, c.U32, c.I32, c.I32, c.I32}
	_typs_PU32I32F32                            = []c.Type{c.Pointer, c.U32, c.I32, c.F32}
	_typs_PVec2PF32F32PU32                      = []c.Type{c.Pointer, _TypVec2, c.Pointer, c.F32, c.F32, c.Pointer, c.U32}
	_typs_PVec2PI32I32PU32                      = []c.Type{c.Pointer, _TypVec2, c.Pointer, c.I32, c.I32, c.Pointer, c.U32}
	_typs_PU32PI32F32PPPU32                     = []c.Type{c.Pointer, c.U32, c.Pointer, c.I32, c.F32, c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_PVec2                                 = []c.Type{c.Pointer, _TypVec2}
	_typs_PVec2I32I32                           = []c.Type{c.Pointer, _TypVec2, c.I32, c.I32}
	_typs_PVec2I32U32                           = []c.Type{c.Pointer, _TypVec2, c.I32, c.U32}
	_typs_PVec2U32                              = []c.Type{c.Pointer, _TypVec2, c.U32}
	_typs_PI32U32                               = []c.Type{c.Pointer, c.I32, c.U32}
	_typs_PVec2Vec2Vec2Vec4Vec4                 = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec4, _TypVec4}
	_typs_U32                                   = []c.Type{c.U32}
	_typs_PU32F32                               = []c.Type{c.Pointer, c.U32, c.F32}
	_typs_U32F32                                = []c.Type{c.U32, c.F32}
	_typs_U32P                                  = []c.Type{c.U32, c.Pointer}
	_typs_U32F32F32                             = []c.Type{c.U32, c.F32, c.F32}
	_typs_U32U32                                = []c.Type{c.U32, c.U32}
	_typs_U32Vec4                               = []c.Type{c.U32, _TypVec4}
	_typs_U32Vec2U32P                           = []c.Type{c.U32, _TypVec2, c.U32, c.Pointer}
	_typs_U32Vec2U32                            = []c.Type{c.U32, _TypVec2, c.U32}
	_typs_Vec2                                  = []c.Type{_TypVec2}
	_typs_I16I16                                = []c.Type{c.I16, c.I16}
	_typs_Vec2U32                               = []c.Type{_TypVec2, c.U32}
	_typs_Vec2U32Vec2                           = []c.Type{_TypVec2, c.U32, _TypVec2}
	_typs_Vec2Vec2PP                            = []c.Type{_TypVec2, _TypVec2, c.Pointer, c.Pointer}
	_typs_Vec2Vec2I32                           = []c.Type{_TypVec2, _TypVec2, c.I32}
	_typs_Vec2Vec2                              = []c.Type{_TypVec2, _TypVec2}
	_typs_Vec4                                  = []c.Type{_TypVec4}
	_typs_I32I32I32I32                          = []c.Type{c.I32, c.I32, c.I32, c.I32}
	_typs_Vec4P                                 = []c.Type{_TypVec4, c.Pointer}
	_typs_PF32F32F32F32                         = []c.Type{c.Pointer, c.F32, c.F32, c.F32, c.F32}
	_typs_PVec2Vec2I32                          = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.I32}
	_typs_PVec2Vec2U32F32                       = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.U32, c.F32}
	_typs_PVec2Vec2U32F32U32F32                 = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.U32, c.F32, c.U32, c.F32}
	_typs_PVec2Vec2U32U32U32U32                 = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.U32, c.U32, c.U32, c.U32}
	_typs_PVec2Vec2Vec2Vec2U32F32               = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32, c.F32}
	_typs_PVec2Vec2U32F32U32                    = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.U32, c.F32, c.U32}
	_typs_PVec2Vec2Vec2Vec2U32                  = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32}
	_typs_PPVec2Vec2Vec2Vec2U32                 = []c.Type{c.Pointer, c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32}
	_typs_PVec2Vec2Vec2U32F32                   = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, c.U32, c.F32}
	_typs_PVec2Vec2Vec2U32                      = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, c.U32}
	_typs_PVec2F32U32I32F32                     = []c.Type{c.Pointer, _TypVec2, c.F32, c.U32, c.I32, c.F32}
	_typs_PVec2F32U32I32                        = []c.Type{c.Pointer, _TypVec2, c.F32, c.U32, c.I32}
	_typs_PVec2U32PP                            = []c.Type{c.Pointer, _TypVec2, c.U32, c.Pointer, c.Pointer}
	_typs_PPF32Vec2U32PPF32P                    = []c.Type{c.Pointer, c.Pointer, c.F32, _TypVec2, c.U32, c.Pointer, c.Pointer, c.F32, c.Pointer}
	_typs_PPI32U32U32F32                        = []c.Type{c.Pointer, c.Pointer, c.I32, c.U32, c.U32, c.F32}
	_typs_PPI32U32                              = []c.Type{c.Pointer, c.Pointer, c.I32, c.U32}
	_typs_PVec2Vec2Vec2Vec2U32F32I32            = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32, c.F32, c.I32}
	_typs_PVec2Vec2Vec2U32F32I32                = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, c.U32, c.F32, c.I32}
	_typs_PPVec2Vec2Vec2Vec2Vec2Vec2Vec2Vec2U32 = []c.Type{c.Pointer, c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32}
	_typs_PPVec2Vec2Vec2Vec2U32F32U32           = []c.Type{c.Pointer, c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32, c.F32, c.U32}
	_typs_PU32U32F32                            = []c.Type{c.Pointer, c.U32, c.U32, c.F32}
	_typs_PVec2F32F32F32I32                     = []c.Type{c.Pointer, _TypVec2, c.F32, c.F32, c.F32, c.I32}
	_typs_PVec2F32I32I32                        = []c.Type{c.Pointer, _TypVec2, c.F32, c.I32, c.I32}
	_typs_PVec2Vec2Vec2I32                      = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, c.I32}
	_typs_PVec2Vec2F32U32                       = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.F32, c.U32}
	_typs_PVec2Vec2U32                          = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.U32}
	_typs_PVec2Vec2Vec2Vec2Vec2Vec2Vec2Vec2U32  = []c.Type{c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.U32}
	_typs_PVec2F32I32I32I32                     = []c.Type{c.Pointer, _TypVec2, c.F32, c.I32, c.I32, c.I32}
	_typs_PPU16I32I32F32Vec2                    = []c.Type{c.Pointer, c.Pointer, c.U16, c.I32, c.I32, c.F32, _TypVec2}
	_typs_PPPP                                  = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer}
	_typs_PI32PPVec2Vec2                        = []c.Type{c.Pointer, c.I32, c.Pointer, c.Pointer, _TypVec2, _TypVec2}
	_typs_PF32PPF32                             = []c.Type{c.Pointer, c.F32, c.Pointer, c.Pointer, c.F32}
	_typs_PPF32Vec2U32U16                       = []c.Type{c.Pointer, c.Pointer, c.F32, _TypVec2, c.U32, c.U16}
	_typs_PPF32Vec2U32Vec4PPF32I32              = []c.Type{c.Pointer, c.Pointer, c.F32, _TypVec2, c.U32, _TypVec4, c.Pointer, c.Pointer, c.F32, c.I32}
	_typs_PPU16F32F32F32F32F32F32F32F32F32      = []c.Type{c.Pointer, c.Pointer, c.U16, c.F32, c.F32, c.F32, c.F32, c.F32, c.F32, c.F32, c.F32, c.F32}
	_typs_PU16U16I32                            = []c.Type{c.Pointer, c.U16, c.U16, c.I32}
	_typs_PU16I32                               = []c.Type{c.Pointer, c.U16, c.I32}
	_typs_PU32U32                               = []c.Type{c.Pointer, c.U32, c.U32}
	_typs_PU64U32                               = []c.Type{c.Pointer, c.U64, c.U32}
	_typs_PRect                                 = []c.Type{c.Pointer, _TypRect}
	_typs_U32Vec2                               = []c.Type{c.U32, _TypVec2}
	_typs_PF32I32                               = []c.Type{c.Pointer, c.F32, c.I32}
	_typs_PPVec2                                = []c.Type{c.Pointer, c.Pointer, _TypVec2}
	_typs_PPVec2Vec2                            = []c.Type{c.Pointer, c.Pointer, _TypVec2, _TypVec2}
	_typs_PPRect                                = []c.Type{c.Pointer, c.Pointer, _TypRect}
	_typs_PVec2Vec2                             = []c.Type{c.Pointer, _TypVec2, _TypVec2}
	_typs_PRectU32                              = []c.Type{c.Pointer, _TypRect, c.U32}
	_typs_PPRectU32                             = []c.Type{c.Pointer, c.Pointer, _TypRect, c.U32}
	_typs_Vec2F32                               = []c.Type{_TypVec2, c.F32}
	_typs_RectF32                               = []c.Type{_TypRect, c.F32}
	_typs_RectU32PU32                           = []c.Type{_TypRect, c.U32, c.Pointer, c.U32}
	_typs_RectU32U32                            = []c.Type{_TypRect, c.U32, c.U32}
	_typs_RectU32                               = []c.Type{_TypRect, c.U32}
	_typs_U32U32U32Rect                         = []c.Type{c.U32, c.U32, c.U32, _TypRect}
	_typs_PVec2F32F32                           = []c.Type{c.Pointer, _TypVec2, c.F32, c.F32}
	_typs_PU32Vec2I32U32                        = []c.Type{c.Pointer, c.U32, _TypVec2, c.I32, c.U32}
	_typs_PVec2Vec2PRectRectU32                 = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.Pointer, _TypRect, _TypRect, c.U32}
	_typs_PPI32F32U32                           = []c.Type{c.Pointer, c.Pointer, c.I32, c.F32, c.U32}
	_typs_U32RectU32                            = []c.Type{c.U32, _TypRect, c.U32}
	_typs_I32I32U32U32                          = []c.Type{c.I32, c.I32, c.U32, c.U32}
	_typs_U32PP                                 = []c.Type{c.U32, c.Pointer, c.Pointer}
	_typs_PU32U32U32U32                         = []c.Type{c.Pointer, c.U32, c.U32, c.U32, c.U32}
	_typs_U32U32U32                             = []c.Type{c.U32, c.U32, c.U32}
	_typs_I32U32U32                             = []c.Type{c.I32, c.U32, c.U32}
	_typs_PPPPI32F32I32                         = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.I32, c.F32, c.I32}
	_typs_PPPPI32I32P                           = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.I32, c.I32, c.Pointer}
	_typs_U32I32F32PP                           = []c.Type{c.U32, c.I32, c.F32, c.Pointer, c.Pointer}
	_typs_U32U32P                               = []c.Type{c.U32, c.U32, c.Pointer}
	_typs_Rect                                  = []c.Type{_TypRect}
	_typs_I32U32I32                             = []c.Type{c.I32, c.U32, c.I32}
	_typs_PU32I32U32Vec2F32                     = []c.Type{c.Pointer, c.U32, c.I32, c.U32, _TypVec2, c.F32}
	_typs_PRectU32P                             = []c.Type{c.Pointer, _TypRect, c.U32, c.Pointer}
	_typs_PPPU32P                               = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.U32, c.Pointer}
	_typs_PRectU32Vec2PU32U32I32PP              = []c.Type{c.Pointer, _TypRect, c.U32, _TypVec2, c.Pointer, c.U32, c.U32, c.I32, c.Pointer, c.Pointer}
	_typs_PRectU32U32                           = []c.Type{c.Pointer, _TypRect, c.U32, c.U32}
	_typs_Vec2PPI32                             = []c.Type{_TypVec2, c.Pointer, c.Pointer, c.I32}
	_typs_Vec2PPF32                             = []c.Type{_TypVec2, c.Pointer, c.Pointer, c.F32}
	_typs_Vec2Vec2PPPVec2P                      = []c.Type{_TypVec2, _TypVec2, c.Pointer, c.Pointer, c.Pointer, _TypVec2, c.Pointer}
	_typs_PVec2Vec2PPPVec2P                     = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.Pointer, c.Pointer, c.Pointer, _TypVec2, c.Pointer}
	_typs_PVec2Vec2F32F32PPP                    = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.F32, c.F32, c.Pointer, c.Pointer, c.Pointer}
	_typs_Vec2Vec2U32I32F32                     = []c.Type{_TypVec2, _TypVec2, c.U32, c.I32, c.F32}
	_typs_Vec2Vec2F32                           = []c.Type{_TypVec2, _TypVec2, c.F32}
	_typs_PVec2Vec2U32F32Vec2F32U32             = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.U32, c.F32, _TypVec2, c.F32, c.U32}
	_typs_Vec2F32I32U32U32U32                   = []c.Type{_TypVec2, c.F32, c.I32, c.U32, c.U32, c.U32}
	_typs_PVec2U32I32F32                        = []c.Type{c.Pointer, _TypVec2, c.U32, c.I32, c.F32}
	_typs_PVec2U32F32                           = []c.Type{c.Pointer, _TypVec2, c.U32, c.F32}
	_typs_PVec2Vec2I32U32                       = []c.Type{c.Pointer, _TypVec2, _TypVec2, c.I32, c.U32}
	_typs_PVec2F32U32                           = []c.Type{c.Pointer, _TypVec2, c.F32, c.U32}
	_typs_PRectU32F32F32F32                     = []c.Type{c.Pointer, _TypRect, c.U32, c.F32, c.F32, c.F32}
	_typs_PRectRectU32F32                       = []c.Type{c.Pointer, _TypRect, _TypRect, c.U32, c.F32}
	_typs_RectRectF32                           = []c.Type{_TypRect, _TypRect, c.F32}
	_typs_PI32Vec2U32                           = []c.Type{c.Pointer, c.I32, _TypVec2, c.U32}
	_typs_U32PVec2Vec2Vec2Vec4Vec4U32           = []c.Type{c.U32, c.Pointer, _TypVec2, _TypVec2, _TypVec2, _TypVec4, _TypVec4, c.U32}
	_typs_U32PPF32                              = []c.Type{c.U32, c.Pointer, c.Pointer, c.F32}
	_typs_PPI64                                 = []c.Type{c.Pointer, c.Pointer, c.I64}
	_typs_PPU64                                 = []c.Type{c.Pointer, c.Pointer, c.U64}
	_typs_U32Vec2P                              = []c.Type{c.U32, _TypVec2, c.Pointer}
	_typs_RectU32I32PI64I64U32                  = []c.Type{_TypRect, c.U32, c.I32, c.Pointer, c.I64, c.I64, c.U32}
	_typs_RectU32PPU32                          = []c.Type{_TypRect, c.U32, c.Pointer, c.Pointer, c.U32}
	_typs_U32U32PF32PPPU32                      = []c.Type{c.U32, c.U32, c.Pointer, c.F32, c.Pointer, c.Pointer, c.Pointer, c.U32}
	_typs_RectU32U32PPPPU32P                    = []c.Type{_TypRect, c.U32, c.U32, c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.U32, c.Pointer}
	_typs_RectU32I32PPF32F32F32F32U32           = []c.Type{_TypRect, c.U32, c.I32, c.Pointer, c.Pointer, c.F32, c.F32, c.F32, c.F32, c.U32}
	_typs_U32U32PP                              = []c.Type{c.U32, c.U32, c.Pointer, c.Pointer}
	_typs_PI32U32PP                             = []c.Type{c.Pointer, c.I32, c.U32, c.Pointer, c.Pointer}
	_typs_U32I32PPP                             = []c.Type{c.U32, c.I32, c.Pointer, c.Pointer, c.Pointer}
	_typs_PU32PP                                = []c.Type{c.Pointer, c.U32, c.Pointer, c.Pointer}
	_typs_U32PPP                                = []c.Type{c.U32, c.Pointer, c.Pointer, c.Pointer}
	_typs_RectU32PPI32U32                       = []c.Type{_TypRect, c.U32, c.Pointer, c.Pointer, c.I32, c.U32}
	_typs_RectU32PU32PPPP                       = []c.Type{_TypRect, c.U32, c.Pointer, c.U32, c.Pointer, c.Pointer, c.Pointer, c.Pointer}
	_typs_PI32I32Vec2Vec2U32U32                 = []c.Type{c.Pointer, c.I32, c.I32, _TypVec2, _TypVec2, c.U32, c.U32}
	_typs_PI32I32Vec2Vec2Vec2Vec2I32            = []c.Type{c.Pointer, c.I32, c.I32, _TypVec2, _TypVec2, _TypVec2, _TypVec2, c.I32}
	_typs_PI32P                                 = []c.Type{c.Pointer, c.I32, c.Pointer}
	_typs_PPPF32F32                             = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.F32, c.F32}
	_typs_PI32I32I32I32PU8U8                    = []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer, c.U8, c.U8}
	_typs_PI32I32I32I32PU8U32                   = []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer, c.U8, c.U32}
)

func _i32ZeroPtr() unsafe.Pointer   { return unsafe.Pointer(&C.__i32_zero) }
func _boolFalsePtr() unsafe.Pointer { return unsafe.Pointer(&C.__i32_false) }
func _u32ZeroPtr() unsafe.Pointer   { return unsafe.Pointer(&C.__u32_zero) }
func _ptrZeroPtr() unsafe.Pointer   { return unsafe.Pointer(&C.__ptr_zero) }
func _vec2ZeroPtr() unsafe.Pointer  { return unsafe.Pointer(&C.__imvec2_zero) }
func _vec2onesPtr() unsafe.Pointer  { return unsafe.Pointer(&C.__imvec2_ones) }
func _vec2_10Ptr() unsafe.Pointer   { return unsafe.Pointer(&C.__imvec2_10) }
func _vec2_01Ptr() unsafe.Pointer   { return unsafe.Pointer(&C.__imvec2_01) }
func _vec4OnesPtr() unsafe.Pointer  { return unsafe.Pointer(&C.__imvec4_ones) }
func _vec4ZeroPtr() unsafe.Pointer  { return unsafe.Pointer(&C.__imvec4_ones) }

var (
	ImVec2ZeroPtr = (*ImVec2)(unsafe.Pointer(&C.__imvec2_zero))
)

var (
	_func_ImGui_ImplOpenGL3_Init_                        = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_Init", OutType: c.I32, InTypes: _typs_P}
	_func_ImGui_ImplOpenGL3_CreateFontsTexture_          = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_CreateFontsTexture", OutType: c.I32, InTypes: nil}
	_func_ImGui_ImplOpenGL3_CreateDeviceObjects_         = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_CreateDeviceObjects", OutType: c.I32, InTypes: nil}
	_func_ImGui_ImplGlfw_InitForOpenGL_                  = &c.FuncPrototype{Name: "ImGui_ImplGlfw_InitForOpenGL", OutType: c.I32, InTypes: _typs_PI32}
	_func_ImGui_ImplOpenGL3_RenderDrawData_              = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_RenderDrawData", OutType: c.Void, InTypes: _typs_P}
	_func_ImGui_ImplOpenGL3_NewFrame_                    = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_NewFrame", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplGlfw_NewFrame_                       = &c.FuncPrototype{Name: "ImGui_ImplGlfw_NewFrame", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplOpenGL3_Shutdown_                    = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_Shutdown", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplOpenGL3_DestroyFontsTexture_         = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_DestroyFontsTexture", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplOpenGL3_DestroyDeviceObjects_        = &c.FuncPrototype{Name: "ImGui_ImplOpenGL3_DestroyDeviceObjects", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplGlfw_WindowFocusCallback_            = &c.FuncPrototype{Name: "ImGui_ImplGlfw_WindowFocusCallback", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImGui_ImplGlfw_Shutdown_                       = &c.FuncPrototype{Name: "ImGui_ImplGlfw_Shutdown", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplGlfw_SetCallbacksChainForAllWindows_ = &c.FuncPrototype{Name: "ImGui_ImplGlfw_SetCallbacksChainForAllWindows", OutType: c.Void, InTypes: _typs_I32}
	_func_ImGui_ImplGlfw_ScrollCallback_                 = &c.FuncPrototype{Name: "ImGui_ImplGlfw_ScrollCallback", OutType: c.Void, InTypes: _typs_PF64F64}
	_func_ImGui_ImplGlfw_RestoreCallbacks_               = &c.FuncPrototype{Name: "ImGui_ImplGlfw_RestoreCallbacks", OutType: c.Void, InTypes: nil}
	_func_ImGui_ImplGlfw_MouseButtonCallback_            = &c.FuncPrototype{Name: "ImGui_ImplGlfw_MouseButtonCallback", OutType: c.Void, InTypes: _typs_PI32I32I32}
	_func_ImGui_ImplGlfw_MonitorCallback_                = &c.FuncPrototype{Name: "ImGui_ImplGlfw_MonitorCallback", OutType: c.Void, InTypes: _typs_PU32}
	_func_ImGui_ImplGlfw_KeyCallback_                    = &c.FuncPrototype{Name: "ImGui_ImplGlfw_KeyCallback", OutType: c.Void, InTypes: _typs_PI32I32I32I32}
	_func_ImGui_ImplGlfw_CursorPosCallback_              = &c.FuncPrototype{Name: "ImGui_ImplGlfw_CursorPosCallback", OutType: c.Void, InTypes: _typs_PF64F64}
	_func_ImGui_ImplGlfw_CursorEnterCallback_            = &c.FuncPrototype{Name: "ImGui_ImplGlfw_CursorEnterCallback", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImGui_ImplGlfw_CharCallback_                   = &c.FuncPrototype{Name: "ImGui_ImplGlfw_CharCallback", OutType: c.Void, InTypes: _typs_PU32}

	_func_ImGuiInputTextCallbackData_ImGuiInputTextCallbackData_ = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_ImGuiInputTextCallbackData", OutType: c.Pointer, InTypes: nil}
	_func_igGetColumnIndex_                                      = &c.FuncPrototype{Name: "igGetColumnIndex", OutType: c.I32, InTypes: nil}
	_func_igGetDrawListSharedData_                               = &c.FuncPrototype{Name: "igGetDrawListSharedData", OutType: c.Pointer, InTypes: nil}
	_func_igTableGetSortSpecs_                                   = &c.FuncPrototype{Name: "igTableGetSortSpecs", OutType: c.Pointer, InTypes: nil}
	_func_igTableGetColumnFlags_                                 = &c.FuncPrototype{Name: "igTableGetColumnFlags", OutType: c.U32, InTypes: _typs_I32}
	_func_igSetNextItemAllowOverlap_                             = &c.FuncPrototype{Name: "igSetNextItemAllowOverlap", OutType: c.Void, InTypes: nil}
	_func_igDestroyPlatformWindows_                              = &c.FuncPrototype{Name: "igDestroyPlatformWindows", OutType: c.Void, InTypes: nil}
	_func_igUpdatePlatformWindows_                               = &c.FuncPrototype{Name: "igUpdatePlatformWindows", OutType: c.Void, InTypes: nil}
	_func_igGetPlatformIO_                                       = &c.FuncPrototype{Name: "igGetPlatformIO", OutType: c.Pointer, InTypes: nil}
	_func_igSetItemDefaultFocus_                                 = &c.FuncPrototype{Name: "igSetItemDefaultFocus", OutType: c.Void, InTypes: nil}
	_func_igGetWindowViewport_                                   = &c.FuncPrototype{Name: "igGetWindowViewport", OutType: c.Pointer, InTypes: nil}
	_func_igGetMainViewport_                                     = &c.FuncPrototype{Name: "igGetMainViewport", OutType: c.Pointer, InTypes: nil}
	_func_igGetMouseCursor_                                      = &c.FuncPrototype{Name: "igGetMouseCursor", OutType: c.I32, InTypes: nil}
	_func_igFindViewportByID_                                    = &c.FuncPrototype{Name: "igFindViewportByID", OutType: c.Pointer, InTypes: _typs_U32}
	_func_igFindViewportByPlatformHandle_                        = &c.FuncPrototype{Name: "igFindViewportByPlatformHandle", OutType: c.Pointer, InTypes: _typs_P}
	_func_igCreateContext_                                       = &c.FuncPrototype{Name: "igCreateContext", OutType: c.Pointer, InTypes: _typs_P}
	_func_igEndDragDropSource_                                   = &c.FuncPrototype{Name: "igEndDragDropSource", OutType: c.Void, InTypes: nil}
	_func_igAcceptDragDropPayload_                               = &c.FuncPrototype{Name: "igAcceptDragDropPayload", OutType: c.Pointer, InTypes: _typs_PU32}
	_func_igEndDragDropTarget_                                   = &c.FuncPrototype{Name: "igEndDragDropTarget", OutType: c.Void, InTypes: nil}
	_func_igGetDragDropPayload_                                  = &c.FuncPrototype{Name: "igGetDragDropPayload", OutType: c.Pointer, InTypes: nil}
	_func_igGetStateStorage_                                     = &c.FuncPrototype{Name: "igGetStateStorage", OutType: c.Pointer, InTypes: nil}
	_func_igGetDrawData_                                         = &c.FuncPrototype{Name: "igGetDrawData", OutType: c.Pointer, InTypes: nil}
	_func_igGetWindowDrawList_                                   = &c.FuncPrototype{Name: "igGetWindowDrawList", OutType: c.Pointer, InTypes: nil}
	_func_igPopButtonRepeat_                                     = &c.FuncPrototype{Name: "igPopButtonRepeat", OutType: c.Void, InTypes: nil}
	_func_igTableHeadersRow_                                     = &c.FuncPrototype{Name: "igTableHeadersRow", OutType: c.Void, InTypes: nil}
	_func_igGetBackgroundDrawList_Nil_                           = &c.FuncPrototype{Name: "igGetBackgroundDrawList_Nil", OutType: c.Pointer, InTypes: nil}
	_func_igGetForegroundDrawList_Nil_                           = &c.FuncPrototype{Name: "igGetForegroundDrawList_Nil", OutType: c.Pointer, InTypes: nil}
	_func_igGetBackgroundDrawList_ViewportPtr_                   = &c.FuncPrototype{Name: "igGetBackgroundDrawList_ViewportPtr", OutType: c.Pointer, InTypes: _typs_P}
	_func_igGetForegroundDrawList_ViewportPtr_                   = &c.FuncPrototype{Name: "igGetForegroundDrawList_ViewportPtr", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImGuiStyle_ImGuiStyle_                                 = &c.FuncPrototype{Name: "ImGuiStyle_ImGuiStyle", OutType: c.Pointer, InTypes: nil}
	_func_igPopTextWrapPos_                                      = &c.FuncPrototype{Name: "igPopTextWrapPos", OutType: c.Void, InTypes: nil}
	_func_igEndChildFrame_                                       = &c.FuncPrototype{Name: "igEndChildFrame", OutType: c.Void, InTypes: nil}
	_func_igGetIO_                                               = &c.FuncPrototype{Name: "igGetIO", OutType: c.Pointer, InTypes: nil}
	_func_igPopItemWidth_                                        = &c.FuncPrototype{Name: "igPopItemWidth", OutType: c.Void, InTypes: nil}
	_func_ImGuiIO_ImGuiIO_                                       = &c.FuncPrototype{Name: "ImGuiIO_ImGuiIO", OutType: c.Pointer, InTypes: nil}
	_func_igEndDisabled_                                         = &c.FuncPrototype{Name: "igEndDisabled", OutType: c.Void, InTypes: nil}
	_func_igPopClipRect_                                         = &c.FuncPrototype{Name: "igPopClipRect", OutType: c.Void, InTypes: nil}
	_func_igPopTabStop_                                          = &c.FuncPrototype{Name: "igPopTabStop", OutType: c.Void, InTypes: nil}
	_func_igBeginGroup_                                          = &c.FuncPrototype{Name: "igBeginGroup", OutType: c.Void, InTypes: nil}
	_func_igEndListBox_                                          = &c.FuncPrototype{Name: "igEndListBox", OutType: c.Void, InTypes: nil}
	_func_igNextColumn_                                          = &c.FuncPrototype{Name: "igNextColumn", OutType: c.Void, InTypes: nil}
	_func_igEndTabItem_                                          = &c.FuncPrototype{Name: "igEndTabItem", OutType: c.Void, InTypes: nil}
	_func_igLogButtons_                                          = &c.FuncPrototype{Name: "igLogButtons", OutType: c.Void, InTypes: nil}
	_func_igSeparator_                                           = &c.FuncPrototype{Name: "igSeparator", OutType: c.Void, InTypes: nil}
	_func_igEndTabBar_                                           = &c.FuncPrototype{Name: "igEndTabBar", OutType: c.Void, InTypes: nil}
	_func_igDockSpace_                                           = &c.FuncPrototype{Name: "igDockSpace", OutType: c.U32, InTypes: _typs_U32Vec2U32P}
	_func_igDockSpaceOverViewport_                               = &c.FuncPrototype{Name: "igDockSpaceOverViewport", OutType: c.U32, InTypes: _typs_PU32P}
	_func_igLogFinish_                                           = &c.FuncPrototype{Name: "igLogFinish", OutType: c.Void, InTypes: nil}
	_func_igGetItemID_                                           = &c.FuncPrototype{Name: "igGetItemID", OutType: c.U32, InTypes: nil}
	_func_igNewFrame_                                            = &c.FuncPrototype{Name: "igNewFrame", OutType: c.Void, InTypes: nil}
	_func_igEndFrame_                                            = &c.FuncPrototype{Name: "igEndFrame", OutType: c.Void, InTypes: nil}
	_func_igEndGroup_                                            = &c.FuncPrototype{Name: "igEndGroup", OutType: c.Void, InTypes: nil}
	_func_igEndCombo_                                            = &c.FuncPrototype{Name: "igEndCombo", OutType: c.Void, InTypes: nil}
	_func_igEndTable_                                            = &c.FuncPrototype{Name: "igEndTable", OutType: c.Void, InTypes: nil}
	_func_igNewLine_                                             = &c.FuncPrototype{Name: "igNewLine", OutType: c.Void, InTypes: nil}
	_func_igSpacing_                                             = &c.FuncPrototype{Name: "igSpacing", OutType: c.Void, InTypes: nil}
	_func_igTreePop_                                             = &c.FuncPrototype{Name: "igTreePop", OutType: c.Void, InTypes: nil}
	_func_igRender_                                              = &c.FuncPrototype{Name: "igRender", OutType: c.Void, InTypes: nil}
	_func_igBullet_                                              = &c.FuncPrototype{Name: "igBullet", OutType: c.Void, InTypes: nil}
	_func_igPopID_                                               = &c.FuncPrototype{Name: "igPopID", OutType: c.Void, InTypes: nil}
	_func_igShowStyleSelector_                                   = &c.FuncPrototype{Name: "igShowStyleSelector", OutType: c.I32, InTypes: _typs_P}
	_func_igGetVersion_                                          = &c.FuncPrototype{Name: "igGetVersion", OutType: c.Pointer, InTypes: nil}
	_func_igBegin_                                               = &c.FuncPrototype{Name: "igBegin", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igBeginChild_Str_                                      = &c.FuncPrototype{Name: "igBeginChild_Str", OutType: c.I32, InTypes: _typs_PVec2I32U32}
	_func_igBeginChild_ID_                                       = &c.FuncPrototype{Name: "igBeginChild_ID", OutType: c.I32, InTypes: _typs_PVec2I32U32}
	_func_igIsWindowAppearing_                                   = &c.FuncPrototype{Name: "igIsWindowAppearing", OutType: c.I32, InTypes: nil}
	_func_igIsWindowCollapsed_                                   = &c.FuncPrototype{Name: "igIsWindowCollapsed", OutType: c.I32, InTypes: nil}
	_func_igIsWindowFocused_                                     = &c.FuncPrototype{Name: "igIsWindowFocused", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsWindowHovered_                                     = &c.FuncPrototype{Name: "igIsWindowHovered", OutType: c.I32, InTypes: _typs_U32}
	_func_igGetWindowDpiScale_                                   = &c.FuncPrototype{Name: "igGetWindowDpiScale", OutType: c.F32, InTypes: nil}
	_func_igGetWindowWidth_                                      = &c.FuncPrototype{Name: "igGetWindowWidth", OutType: c.F32, InTypes: nil}
	_func_igGetWindowHeight_                                     = &c.FuncPrototype{Name: "igGetWindowHeight", OutType: c.F32, InTypes: nil}
	_func_igGetScrollX_                                          = &c.FuncPrototype{Name: "igGetScrollX", OutType: c.F32, InTypes: nil}
	_func_igGetScrollY_                                          = &c.FuncPrototype{Name: "igGetScrollY", OutType: c.F32, InTypes: nil}
	_func_igGetScrollMaxX_                                       = &c.FuncPrototype{Name: "igGetScrollMaxX", OutType: c.F32, InTypes: nil}
	_func_igGetScrollMaxY_                                       = &c.FuncPrototype{Name: "igGetScrollMaxY", OutType: c.F32, InTypes: nil}
	_func_igCalcItemWidth_                                       = &c.FuncPrototype{Name: "igCalcItemWidth", OutType: c.F32, InTypes: nil}
	_func_GetFontSize_                                           = &c.FuncPrototype{Name: "GetFontSize", OutType: c.F32, InTypes: nil}
	_func_igGetColorU32_Col_                                     = &c.FuncPrototype{Name: "igGetColorU32_Col", OutType: c.U32, InTypes: _typs_U32F32}
	_func_igGetColorU32_Vec4_                                    = &c.FuncPrototype{Name: "igGetColorU32_Vec4", OutType: c.U32, InTypes: _typs_Vec4}
	_func_igGetColorU32_U32_                                     = &c.FuncPrototype{Name: "igGetColorU32_U32", OutType: c.U32, InTypes: _typs_U32}
	_func_igGetCursorPosX_                                       = &c.FuncPrototype{Name: "igGetCursorPosX", OutType: c.F32, InTypes: nil}
	_func_igGetCursorPosY_                                       = &c.FuncPrototype{Name: "igGetCursorPosY", OutType: c.F32, InTypes: nil}
	_func_igGetTextLineHeight_                                   = &c.FuncPrototype{Name: "igGetTextLineHeight", OutType: c.F32, InTypes: nil}
	_func_igGetTextLineHeightWithSpacing_                        = &c.FuncPrototype{Name: "igGetTextLineHeightWithSpacing", OutType: c.F32, InTypes: nil}
	_func_igGetFrameHeight_                                      = &c.FuncPrototype{Name: "igGetFrameHeight", OutType: c.F32, InTypes: nil}
	_func_igGetFrameHeightWithSpacing_                           = &c.FuncPrototype{Name: "igGetFrameHeightWithSpacing", OutType: c.F32, InTypes: nil}
	_func_igButton_                                              = &c.FuncPrototype{Name: "igButton", OutType: c.I32, InTypes: _typs_PVec2}
	_func_igSmallButton_                                         = &c.FuncPrototype{Name: "igSmallButton", OutType: c.I32, InTypes: _typs_P}
	_func_igInvisibleButton_                                     = &c.FuncPrototype{Name: "igInvisibleButton", OutType: c.I32, InTypes: _typs_PVec2U32}
	_func_igArrowButton_                                         = &c.FuncPrototype{Name: "igArrowButton", OutType: c.I32, InTypes: _typs_PI32}
	_func_igCheckboxFlags_IntPtr_                                = &c.FuncPrototype{Name: "igCheckboxFlags_IntPtr", OutType: c.I32, InTypes: _typs_PPI32}
	_func_igCheckboxFlags_UintPtr_                               = &c.FuncPrototype{Name: "igCheckboxFlags_UintPtr", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igRadioButton_Bool_                                    = &c.FuncPrototype{Name: "igRadioButton_Bool", OutType: c.I32, InTypes: _typs_PI32}
	_func_igRadioButton_IntPtr_                                  = &c.FuncPrototype{Name: "igRadioButton_IntPtr", OutType: c.I32, InTypes: _typs_PPI32}
	_func_igImageButton_                                         = &c.FuncPrototype{Name: "igImageButton", OutType: c.I32, InTypes: _typs_PPVec2Vec2Vec2Vec4Vec4}
	_func_igBeginCombo_                                          = &c.FuncPrototype{Name: "igBeginCombo", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igCombo_Str_arr_                                       = &c.FuncPrototype{Name: "igCombo_Str_arr", OutType: c.I32, InTypes: _typs_PPPI32I32}
	_func_igCombo_Str_                                           = &c.FuncPrototype{Name: "igCombo_Str", OutType: c.I32, InTypes: _typs_PPPI32}
	_func_igDragFloat_                                           = &c.FuncPrototype{Name: "igDragFloat", OutType: c.I32, InTypes: _typs_PPF32F32F32PU32}
	_func_igDragFloat2_                                          = &c.FuncPrototype{Name: "igDragFloat2", OutType: c.I32, InTypes: _typs_PPF32F32F32PU32}
	_func_igDragFloat3_                                          = &c.FuncPrototype{Name: "igDragFloat3", OutType: c.I32, InTypes: _typs_PPF32F32F32PU32}
	_func_igDragFloat4_                                          = &c.FuncPrototype{Name: "igDragFloat4", OutType: c.I32, InTypes: _typs_PPF32F32F32PU32}
	_func_igDragFloatRange2_                                     = &c.FuncPrototype{Name: "igDragFloatRange2", OutType: c.I32, InTypes: _typs_PPPF32F32F32PPU32}
	_func_igDragInt_                                             = &c.FuncPrototype{Name: "igDragInt", OutType: c.I32, InTypes: _typs_PPF32I32I32PU32}
	_func_igDragInt2_                                            = &c.FuncPrototype{Name: "igDragInt2", OutType: c.I32, InTypes: _typs_PPF32I32I32PU32}
	_func_igDragInt3_                                            = &c.FuncPrototype{Name: "igDragInt3", OutType: c.I32, InTypes: _typs_PPF32I32I32PU32}
	_func_igDragInt4_                                            = &c.FuncPrototype{Name: "igDragInt4", OutType: c.I32, InTypes: _typs_PPF32I32I32PU32}
	_func_igDragIntRange2_                                       = &c.FuncPrototype{Name: "igDragIntRange2", OutType: c.I32, InTypes: _typs_PPPF32I32I32PPU32}
	_func_igDragScalar_                                          = &c.FuncPrototype{Name: "igDragScalar", OutType: c.I32, InTypes: _typs_PU32PF32PPPU32}
	_func_igDragScalarN_                                         = &c.FuncPrototype{Name: "igDragScalarN", OutType: c.I32, InTypes: _typs_PU32PI32F32PPPU32}
	_func_igSliderFloat_                                         = &c.FuncPrototype{Name: "igSliderFloat", OutType: c.I32, InTypes: _typs_PPF32F32PU32}
	_func_igSliderFloat2_                                        = &c.FuncPrototype{Name: "igSliderFloat2", OutType: c.I32, InTypes: _typs_PPF32F32PU32}
	_func_igSliderFloat3_                                        = &c.FuncPrototype{Name: "igSliderFloat3", OutType: c.I32, InTypes: _typs_PPF32F32PU32}
	_func_igSliderFloat4_                                        = &c.FuncPrototype{Name: "igSliderFloat4", OutType: c.I32, InTypes: _typs_PPF32F32PU32}
	_func_igSliderAngle_                                         = &c.FuncPrototype{Name: "igSliderAngle", OutType: c.I32, InTypes: _typs_PPF32F32PU32}
	_func_igSliderInt_                                           = &c.FuncPrototype{Name: "igSliderInt", OutType: c.I32, InTypes: _typs_PPI32I32PU32}
	_func_igSliderInt2_                                          = &c.FuncPrototype{Name: "igSliderInt2", OutType: c.I32, InTypes: _typs_PPI32I32PU32}
	_func_igSliderInt3_                                          = &c.FuncPrototype{Name: "igSliderInt3", OutType: c.I32, InTypes: _typs_PPI32I32PU32}
	_func_igSliderInt4_                                          = &c.FuncPrototype{Name: "igSliderInt4", OutType: c.I32, InTypes: _typs_PPI32I32PU32}
	_func_igSliderScalar_                                        = &c.FuncPrototype{Name: "igSliderScalar", OutType: c.I32, InTypes: _typs_PU32PPPPU32}
	_func_igSliderScalarN_                                       = &c.FuncPrototype{Name: "igSliderScalarN", OutType: c.I32, InTypes: _typs_PU32PI32PPPU32}
	_func_igVSliderFloat_                                        = &c.FuncPrototype{Name: "igVSliderFloat", OutType: c.I32, InTypes: _typs_PVec2PF32F32PU32}
	_func_igVSliderInt_                                          = &c.FuncPrototype{Name: "igVSliderInt", OutType: c.I32, InTypes: _typs_PVec2PI32I32PU32}
	_func_igVSliderScalar_                                       = &c.FuncPrototype{Name: "igVSliderScalar", OutType: c.I32, InTypes: _typs_PVec2U32PPPPU32}
	_func_igInputFloat_                                          = &c.FuncPrototype{Name: "igInputFloat", OutType: c.I32, InTypes: _typs_PPF32F32PU32}
	_func_igInputFloat2_                                         = &c.FuncPrototype{Name: "igInputFloat2", OutType: c.I32, InTypes: _typs_PPPU32}
	_func_igInputFloat3_                                         = &c.FuncPrototype{Name: "igInputFloat3", OutType: c.I32, InTypes: _typs_PPPU32}
	_func_igInputFloat4_                                         = &c.FuncPrototype{Name: "igInputFloat4", OutType: c.I32, InTypes: _typs_PPPU32}
	_func_igInputInt_                                            = &c.FuncPrototype{Name: "igInputInt", OutType: c.I32, InTypes: _typs_PPI32I32U32}
	_func_igInputInt2_                                           = &c.FuncPrototype{Name: "igInputInt2", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igInputInt3_                                           = &c.FuncPrototype{Name: "igInputInt3", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igInputInt4_                                           = &c.FuncPrototype{Name: "igInputInt4", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igInputDouble_                                         = &c.FuncPrototype{Name: "igInputDouble", OutType: c.I32, InTypes: _typs_PPF64F64PU32}
	_func_igInputScalar_                                         = &c.FuncPrototype{Name: "igInputScalar", OutType: c.I32, InTypes: _typs_PU32PPPPU32}
	_func_igInputScalarN_                                        = &c.FuncPrototype{Name: "igInputScalarN", OutType: c.I32, InTypes: _typs_PU32PI32PPPU32}
	_func_igColorEdit3_                                          = &c.FuncPrototype{Name: "igColorEdit3", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igColorEdit4_                                          = &c.FuncPrototype{Name: "igColorEdit4", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igColorPicker3_                                        = &c.FuncPrototype{Name: "igColorPicker3", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igColorPicker4_                                        = &c.FuncPrototype{Name: "igColorPicker4", OutType: c.I32, InTypes: _typs_PPU32P}
	_func_igColorButton_                                         = &c.FuncPrototype{Name: "igColorButton", OutType: c.I32, InTypes: _typs_PVec4U32Vec2}
	_func_igTreeNode_Str_                                        = &c.FuncPrototype{Name: "igTreeNode_Str", OutType: c.I32, InTypes: _typs_P}
	_func_igTreeNode_StrStr_                                     = &c.FuncPrototype{Name: "igTreeNode_StrStr", OutType: c.I32, InTypes: _typs_PP}
	_func_igTreeNode_Ptr_                                        = &c.FuncPrototype{Name: "igTreeNode_Ptr", OutType: c.I32, InTypes: _typs_PP}
	_func_igTreeNodeEx_Str_                                      = &c.FuncPrototype{Name: "igTreeNodeEx_Str", OutType: c.I32, InTypes: _typs_PU32}
	_func_igTreeNodeEx_StrStr_                                   = &c.FuncPrototype{Name: "igTreeNodeEx_StrStr", OutType: c.I32, InTypes: _typs_PU32P}
	_func_igTreeNodeEx_Ptr_                                      = &c.FuncPrototype{Name: "igTreeNodeEx_Ptr", OutType: c.I32, InTypes: _typs_PU32P}
	_func_igGetTreeNodeToLabelSpacing_                           = &c.FuncPrototype{Name: "igGetTreeNodeToLabelSpacing", OutType: c.F32, InTypes: nil}
	_func_igCollapsingHeader_TreeNodeFlags_                      = &c.FuncPrototype{Name: "igCollapsingHeader_TreeNodeFlags", OutType: c.I32, InTypes: _typs_PU32}
	_func_igSelectable_Bool_                                     = &c.FuncPrototype{Name: "igSelectable_Bool", OutType: c.I32, InTypes: _typs_PI32U32Vec2}
	_func_igBeginListBox_                                        = &c.FuncPrototype{Name: "igBeginListBox", OutType: c.I32, InTypes: _typs_PVec2}
	_func_igListBox_Str_arr_                                     = &c.FuncPrototype{Name: "igListBox_Str_arr", OutType: c.I32, InTypes: _typs_PPPI32I32}
	_func_igListBox_FnBoolPtr_                                   = &c.FuncPrototype{Name: "igListBox_FnBoolPtr", OutType: c.I32, InTypes: _typs_PPPPI32I32}
	_func_igBeginMenuBar_                                        = &c.FuncPrototype{Name: "igBeginMenuBar", OutType: c.I32, InTypes: nil}
	_func_igBeginMainMenuBar_                                    = &c.FuncPrototype{Name: "igBeginMainMenuBar", OutType: c.I32, InTypes: nil}
	_func_igBeginMenu_                                           = &c.FuncPrototype{Name: "igBeginMenu", OutType: c.I32, InTypes: _typs_PI32}
	_func_igMenuItem_Bool_                                       = &c.FuncPrototype{Name: "igMenuItem_Bool", OutType: c.I32, InTypes: _typs_PPI32I32}
	_func_igBeginTooltip_                                        = &c.FuncPrototype{Name: "igBeginTooltip", OutType: c.I32, InTypes: nil}
	_func_igBeginItemTooltip_                                    = &c.FuncPrototype{Name: "igBeginItemTooltip", OutType: c.I32, InTypes: nil}
	_func_igBeginPopup_                                          = &c.FuncPrototype{Name: "igBeginPopup", OutType: c.I32, InTypes: _typs_PU32}
	_func_igBeginPopupContextItem_                               = &c.FuncPrototype{Name: "igBeginPopupContextItem", OutType: c.I32, InTypes: _typs_PU32}
	_func_igBeginPopupContextWindow_                             = &c.FuncPrototype{Name: "igBeginPopupContextWindow", OutType: c.I32, InTypes: _typs_PU32}
	_func_igBeginPopupContextVoid_                               = &c.FuncPrototype{Name: "igBeginPopupContextVoid", OutType: c.I32, InTypes: _typs_PU32}
	_func_igIsPopupOpen_Str_                                     = &c.FuncPrototype{Name: "igIsPopupOpen_Str", OutType: c.I32, InTypes: _typs_PU32}
	_func_igBeginTable_                                          = &c.FuncPrototype{Name: "igBeginTable", OutType: c.I32, InTypes: _typs_PI32U32Vec2F32}
	_func_igTableNextColumn_                                     = &c.FuncPrototype{Name: "igTableNextColumn", OutType: c.I32, InTypes: nil}
	_func_igTableSetColumnIndex_                                 = &c.FuncPrototype{Name: "igTableSetColumnIndex", OutType: c.I32, InTypes: _typs_I32}
	_func_igTableGetColumnCount_                                 = &c.FuncPrototype{Name: "igTableGetColumnCount", OutType: c.I32, InTypes: nil}
	_func_igTableGetColumnIndex_                                 = &c.FuncPrototype{Name: "igTableGetColumnIndex", OutType: c.I32, InTypes: nil}
	_func_igTableGetRowIndex_                                    = &c.FuncPrototype{Name: "igTableGetRowIndex", OutType: c.I32, InTypes: nil}
	_func_igTableGetColumnName_Int_                              = &c.FuncPrototype{Name: "igTableGetColumnName_Int", OutType: c.Pointer, InTypes: _typs_I32}
	_func_igGetColumnWidth_                                      = &c.FuncPrototype{Name: "igGetColumnWidth", OutType: c.F32, InTypes: _typs_I32}
	_func_igGetColumnOffset_                                     = &c.FuncPrototype{Name: "igGetColumnOffset", OutType: c.F32, InTypes: _typs_I32}
	_func_igGetColumnsCount_                                     = &c.FuncPrototype{Name: "igGetColumnsCount", OutType: c.I32, InTypes: nil}
	_func_igBeginTabBar_                                         = &c.FuncPrototype{Name: "igBeginTabBar", OutType: c.I32, InTypes: _typs_PU32}
	_func_igBeginTabItem_                                        = &c.FuncPrototype{Name: "igBeginTabItem", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igTabItemButton_                                       = &c.FuncPrototype{Name: "igTabItemButton", OutType: c.I32, InTypes: _typs_PU32}
	_func_igIsWindowDocked_                                      = &c.FuncPrototype{Name: "igIsWindowDocked", OutType: c.I32, InTypes: nil}
	_func_igBeginDragDropSource_                                 = &c.FuncPrototype{Name: "igBeginDragDropSource", OutType: c.I32, InTypes: _typs_U32}
	_func_igSetDragDropPayload_                                  = &c.FuncPrototype{Name: "igSetDragDropPayload", OutType: c.I32, InTypes: _typs_PPU64U32}
	_func_igBeginDragDropTarget_                                 = &c.FuncPrototype{Name: "igBeginDragDropTarget", OutType: c.I32, InTypes: nil}
	_func_igIsItemHovered_                                       = &c.FuncPrototype{Name: "igIsItemHovered", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsItemActive_                                        = &c.FuncPrototype{Name: "igIsItemActive", OutType: c.I32, InTypes: nil}
	_func_igIsItemFocused_                                       = &c.FuncPrototype{Name: "igIsItemFocused", OutType: c.I32, InTypes: nil}
	_func_igIsItemClicked_                                       = &c.FuncPrototype{Name: "igIsItemClicked", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsItemVisible_                                       = &c.FuncPrototype{Name: "igIsItemVisible", OutType: c.I32, InTypes: nil}
	_func_igIsItemEdited_                                        = &c.FuncPrototype{Name: "igIsItemEdited", OutType: c.I32, InTypes: nil}
	_func_igIsItemActivated_                                     = &c.FuncPrototype{Name: "igIsItemActivated", OutType: c.I32, InTypes: nil}
	_func_igIsItemDeactivated_                                   = &c.FuncPrototype{Name: "igIsItemDeactivated", OutType: c.I32, InTypes: nil}
	_func_igIsItemDeactivatedAfterEdit_                          = &c.FuncPrototype{Name: "igIsItemDeactivatedAfterEdit", OutType: c.I32, InTypes: nil}
	_func_igIsItemToggledOpen_                                   = &c.FuncPrototype{Name: "igIsItemToggledOpen", OutType: c.I32, InTypes: nil}
	_func_igIsAnyItemHovered_                                    = &c.FuncPrototype{Name: "igIsAnyItemHovered", OutType: c.I32, InTypes: nil}
	_func_igIsAnyItemActive_                                     = &c.FuncPrototype{Name: "igIsAnyItemActive", OutType: c.I32, InTypes: nil}
	_func_igIsAnyItemFocused_                                    = &c.FuncPrototype{Name: "igIsAnyItemFocused", OutType: c.I32, InTypes: nil}
	_func_igIsRectVisible_Nil_                                   = &c.FuncPrototype{Name: "igIsRectVisible_Nil", OutType: c.I32, InTypes: _typs_Vec2}
	_func_igIsRectVisible_Vec2_                                  = &c.FuncPrototype{Name: "igIsRectVisible_Vec2", OutType: c.I32, InTypes: _typs_Vec2Vec2}
	_func_igGetTime_                                             = &c.FuncPrototype{Name: "igGetTime", OutType: c.F64, InTypes: nil}
	_func_igGetFrameCount_                                       = &c.FuncPrototype{Name: "igGetFrameCount", OutType: c.I32, InTypes: nil}
	_func_igGetStyleColorName_                                   = &c.FuncPrototype{Name: "igGetStyleColorName", OutType: c.Pointer, InTypes: _typs_U32}
	_func_igBeginChildFrame_                                     = &c.FuncPrototype{Name: "igBeginChildFrame", OutType: c.I32, InTypes: _typs_U32Vec2U32}
	_func_igColorConvertFloat4ToU32_                             = &c.FuncPrototype{Name: "igColorConvertFloat4ToU32", OutType: c.U32, InTypes: _typs_Vec4}
	_func_igIsKeyDown_Nil_                                       = &c.FuncPrototype{Name: "igIsKeyDown_Nil", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsKeyPressed_Bool_                                   = &c.FuncPrototype{Name: "igIsKeyPressed_Bool", OutType: c.I32, InTypes: _typs_U32I32}
	_func_igIsKeyReleased_Nil_                                   = &c.FuncPrototype{Name: "igIsKeyReleased_Nil", OutType: c.I32, InTypes: _typs_U32}
	_func_igGetKeyPressedAmount_                                 = &c.FuncPrototype{Name: "igGetKeyPressedAmount", OutType: c.I32, InTypes: _typs_U32F32F32}
	_func_igGetKeyName_                                          = &c.FuncPrototype{Name: "igGetKeyName", OutType: c.Pointer, InTypes: _typs_U32}
	_func_igIsMouseDown_Nil_                                     = &c.FuncPrototype{Name: "igIsMouseDown_Nil", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsMouseClicked_Bool_                                 = &c.FuncPrototype{Name: "igIsMouseClicked_Bool", OutType: c.I32, InTypes: _typs_U32I32}
	_func_igIsMouseReleased_Nil_                                 = &c.FuncPrototype{Name: "igIsMouseReleased_Nil", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsMouseDoubleClicked_                                = &c.FuncPrototype{Name: "igIsMouseDoubleClicked", OutType: c.I32, InTypes: _typs_U32}
	_func_igGetMouseClickedCount_                                = &c.FuncPrototype{Name: "igGetMouseClickedCount", OutType: c.I32, InTypes: _typs_U32}
	_func_igIsMouseHoveringRect_                                 = &c.FuncPrototype{Name: "igIsMouseHoveringRect", OutType: c.I32, InTypes: _typs_Vec2Vec2I32}
	_func_igIsMousePosValid_                                     = &c.FuncPrototype{Name: "igIsMousePosValid", OutType: c.I32, InTypes: _typs_P}
	_func_igIsAnyMouseDown_                                      = &c.FuncPrototype{Name: "igIsAnyMouseDown", OutType: c.I32, InTypes: nil}
	_func_igIsMouseDragging_                                     = &c.FuncPrototype{Name: "igIsMouseDragging", OutType: c.I32, InTypes: _typs_U32F32}
	_func_igGetClipboardText_                                    = &c.FuncPrototype{Name: "igGetClipboardText", OutType: c.Pointer, InTypes: nil}
	_func_igDebugCheckVersionAndDataLayout_                      = &c.FuncPrototype{Name: "igDebugCheckVersionAndDataLayout", OutType: c.I32, InTypes: _typs_PU64U64U64U64U64U64}
	_func_igInputText_                                           = &c.FuncPrototype{Name: "igInputText", OutType: c.I32, InTypes: _typs_PPU64U32PP}
	_func_igInputTextMultiline_                                  = &c.FuncPrototype{Name: "igInputTextMultiline", OutType: c.I32, InTypes: _typs_PPU64Vec2U32PP}
	_func_igInputTextWithHint_                                   = &c.FuncPrototype{Name: "igInputTextWithHint", OutType: c.I32, InTypes: _typs_PPPU64U32PP}
	_func_ImVec2_ImVec2_Nil_                                     = &c.FuncPrototype{Name: "ImVec2_ImVec2_Nil", OutType: c.Pointer, InTypes: nil}
	_func_ImVec2_ImVec2_Float_                                   = &c.FuncPrototype{Name: "ImVec2_ImVec2_Float", OutType: c.Pointer, InTypes: _typs_F32F32}
	_func_ImVec4_ImVec4_Nil_                                     = &c.FuncPrototype{Name: "ImVec4_ImVec4_Nil", OutType: c.Pointer, InTypes: nil}
	_func_ImVec4_ImVec4_Float_                                   = &c.FuncPrototype{Name: "ImVec4_ImVec4_Float", OutType: c.Pointer, InTypes: _typs_F32F32F32F32}
	_func_igGetCurrentContext_                                   = &c.FuncPrototype{Name: "igGetCurrentContext", OutType: c.Pointer, InTypes: nil}
	_func_igGetStyle_                                            = &c.FuncPrototype{Name: "igGetStyle", OutType: c.Pointer, InTypes: nil}
	_func_igGetFont_                                             = &c.FuncPrototype{Name: "igGetFont", OutType: c.Pointer, InTypes: nil}
	_func_igGetStyleColorVec4_                                   = &c.FuncPrototype{Name: "igGetStyleColorVec4", OutType: c.Pointer, InTypes: _typs_U32}
	_func_igGetID_Str_                                           = &c.FuncPrototype{Name: "igGetID_Str", OutType: c.U32, InTypes: _typs_P}
	_func_igGetID_StrStr_                                        = &c.FuncPrototype{Name: "igGetID_StrStr", OutType: c.U32, InTypes: _typs_PP}
	_func_igGetID_Ptr_                                           = &c.FuncPrototype{Name: "igGetID_Ptr", OutType: c.U32, InTypes: _typs_P}
	_func_igCheckbox_                                            = &c.FuncPrototype{Name: "igCheckbox", OutType: c.I32, InTypes: _typs_PP}
	_func_igCollapsingHeader_BoolPtr_                            = &c.FuncPrototype{Name: "igCollapsingHeader_BoolPtr", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igSelectable_BoolPtr_                                  = &c.FuncPrototype{Name: "igSelectable_BoolPtr", OutType: c.I32, InTypes: _typs_PPU32Vec2}
	_func_igMenuItem_BoolPtr_                                    = &c.FuncPrototype{Name: "igMenuItem_BoolPtr", OutType: c.I32, InTypes: _typs_PPPI32}
	_func_igBeginPopupModal_                                     = &c.FuncPrototype{Name: "igBeginPopupModal", OutType: c.I32, InTypes: _typs_PPU32}
	_func_igSaveIniSettingsToMemory_                             = &c.FuncPrototype{Name: "igSaveIniSettingsToMemory", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontConfig_ImFontConfig_                             = &c.FuncPrototype{Name: "ImFontConfig_ImFontConfig", OutType: c.Pointer, InTypes: nil}
	_func_ImFontAtlas_ImFontAtlas_                               = &c.FuncPrototype{Name: "ImFontAtlas_ImFontAtlas", OutType: c.Pointer, InTypes: nil}
	_func_ImFontAtlas_AddFont_                                   = &c.FuncPrototype{Name: "ImFontAtlas_AddFont", OutType: c.Pointer, InTypes: _typs_PP}
	_func_ImFontAtlas_AddFontDefault_                            = &c.FuncPrototype{Name: "ImFontAtlas_AddFontDefault", OutType: c.Pointer, InTypes: _typs_PP}
	_func_ImFontAtlas_AddFontFromFileTTF_                        = &c.FuncPrototype{Name: "ImFontAtlas_AddFontFromFileTTF", OutType: c.Pointer, InTypes: _typs_PPF32PP}
	_func_ImFontAtlas_AddFontFromMemoryTTF_                      = &c.FuncPrototype{Name: "ImFontAtlas_AddFontFromMemoryTTF", OutType: c.Pointer, InTypes: _typs_PPI32F32PP}
	_func_ImFontAtlas_AddFontFromMemoryCompressedTTF_            = &c.FuncPrototype{Name: "ImFontAtlas_AddFontFromMemoryCompressedTTF", OutType: c.Pointer, InTypes: _typs_PPI32F32PP}
	_func_ImFontAtlas_AddFontFromMemoryCompressedBase85TTF_      = &c.FuncPrototype{Name: "ImFontAtlas_AddFontFromMemoryCompressedBase85TTF", OutType: c.Pointer, InTypes: _typs_PPF32PP}
	_func_ImFontAtlas_GetGlyphRangesDefault_                     = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesDefault", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesGreek_                       = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesGreek", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesKorean_                      = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesKorean", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesJapanese_                    = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesJapanese", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesChineseFull_                 = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesChineseFull", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon_     = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesCyrillic_                    = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesCyrillic", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesThai_                        = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesThai", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFontAtlas_GetGlyphRangesVietnamese_                  = &c.FuncPrototype{Name: "ImFontAtlas_GetGlyphRangesVietnamese", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImVec2_destroy_                                        = &c.FuncPrototype{Name: "ImVec2_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImVec4_destroy_                                        = &c.FuncPrototype{Name: "ImVec4_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_igDestroyContext_                                      = &c.FuncPrototype{Name: "igDestroyContext", OutType: c.Void, InTypes: _typs_P}
	_func_igSetCurrentContext_                                   = &c.FuncPrototype{Name: "igSetCurrentContext", OutType: c.Void, InTypes: _typs_P}
	_func_igShowDemoWindow_                                      = &c.FuncPrototype{Name: "igShowDemoWindow", OutType: c.Void, InTypes: _typs_P}
	_func_igShowMetricsWindow_                                   = &c.FuncPrototype{Name: "igShowMetricsWindow", OutType: c.Void, InTypes: _typs_P}
	_func_igShowDebugLogWindow_                                  = &c.FuncPrototype{Name: "igShowDebugLogWindow", OutType: c.Void, InTypes: _typs_P}
	_func_igShowStackToolWindow_                                 = &c.FuncPrototype{Name: "igShowStackToolWindow", OutType: c.Void, InTypes: _typs_P}
	_func_igShowAboutWindow_                                     = &c.FuncPrototype{Name: "igShowAboutWindow", OutType: c.Void, InTypes: _typs_P}
	_func_igShowStyleEditor_                                     = &c.FuncPrototype{Name: "igShowStyleEditor", OutType: c.Void, InTypes: _typs_P}
	_func_igShowFontSelector_                                    = &c.FuncPrototype{Name: "igShowFontSelector", OutType: c.Void, InTypes: _typs_P}
	_func_igShowUserGuide_                                       = &c.FuncPrototype{Name: "igShowUserGuide", OutType: c.Void, InTypes: nil}
	_func_igStyleColorsDark_                                     = &c.FuncPrototype{Name: "igStyleColorsDark", OutType: c.Void, InTypes: _typs_P}
	_func_igStyleColorsLight_                                    = &c.FuncPrototype{Name: "igStyleColorsLight", OutType: c.Void, InTypes: _typs_P}
	_func_igStyleColorsClassic_                                  = &c.FuncPrototype{Name: "igStyleColorsClassic", OutType: c.Void, InTypes: _typs_P}
	_func_igEnd_                                                 = &c.FuncPrototype{Name: "igEnd", OutType: c.Void, InTypes: nil}
	_func_igEndChild_                                            = &c.FuncPrototype{Name: "igEndChild", OutType: c.Void, InTypes: nil}
	_func_igGetWindowPos_                                        = &c.FuncPrototype{Name: "igGetWindowPos", OutType: c.Void, InTypes: _typs_P}
	_func_igGetWindowSize_                                       = &c.FuncPrototype{Name: "igGetWindowSize", OutType: c.Void, InTypes: _typs_P}
	_func_igSetNextWindowPos_                                    = &c.FuncPrototype{Name: "igSetNextWindowPos", OutType: c.Void, InTypes: _typs_Vec2U32Vec2}
	_func_igSetNextWindowSize_                                   = &c.FuncPrototype{Name: "igSetNextWindowSize", OutType: c.Void, InTypes: _typs_Vec2U32}
	_func_igSetNextWindowSizeConstraints_                        = &c.FuncPrototype{Name: "igSetNextWindowSizeConstraints", OutType: c.Void, InTypes: _typs_Vec2Vec2PP}
	_func_igSetNextWindowContentSize_                            = &c.FuncPrototype{Name: "igSetNextWindowContentSize", OutType: c.Void, InTypes: _typs_Vec2}
	_func_igSetNextWindowCollapsed_                              = &c.FuncPrototype{Name: "igSetNextWindowCollapsed", OutType: c.Void, InTypes: _typs_I32U32}
	_func_igSetNextWindowFocus_                                  = &c.FuncPrototype{Name: "igSetNextWindowFocus", OutType: c.Void, InTypes: nil}
	_func_igSetNextWindowScroll_                                 = &c.FuncPrototype{Name: "igSetNextWindowScroll", OutType: c.Void, InTypes: _typs_Vec2}
	_func_igSetNextWindowBgAlpha_                                = &c.FuncPrototype{Name: "igSetNextWindowBgAlpha", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetNextWindowViewport_                               = &c.FuncPrototype{Name: "igSetNextWindowViewport", OutType: c.Void, InTypes: _typs_U32}
	_func_igSetWindowPos_Vec2_                                   = &c.FuncPrototype{Name: "igSetWindowPos_Vec2", OutType: c.Void, InTypes: _typs_Vec2U32}
	_func_igSetWindowSize_Vec2_                                  = &c.FuncPrototype{Name: "igSetWindowSize_Vec2", OutType: c.Void, InTypes: _typs_Vec2U32}
	_func_igSetWindowCollapsed_Bool_                             = &c.FuncPrototype{Name: "igSetWindowCollapsed_Bool", OutType: c.Void, InTypes: _typs_I32U32}
	_func_igSetWindowFocus_Nil_                                  = &c.FuncPrototype{Name: "igSetWindowFocus_Nil", OutType: c.Void, InTypes: nil}
	_func_igSetWindowFontScale_                                  = &c.FuncPrototype{Name: "igSetWindowFontScale", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetWindowPos_Str_                                    = &c.FuncPrototype{Name: "igSetWindowPos_Str", OutType: c.Void, InTypes: _typs_PVec2U32}
	_func_igSetWindowSize_Str_                                   = &c.FuncPrototype{Name: "igSetWindowSize_Str", OutType: c.Void, InTypes: _typs_PVec2U32}
	_func_igSetWindowCollapsed_Str_                              = &c.FuncPrototype{Name: "igSetWindowCollapsed_Str", OutType: c.Void, InTypes: _typs_PI32U32}
	_func_igSetWindowFocus_Str_                                  = &c.FuncPrototype{Name: "igSetWindowFocus_Str", OutType: c.Void, InTypes: _typs_P}
	_func_igGetContentRegionAvail_                               = &c.FuncPrototype{Name: "igGetContentRegionAvail", OutType: c.Void, InTypes: _typs_P}
	_func_igGetContentRegionMax_                                 = &c.FuncPrototype{Name: "igGetContentRegionMax", OutType: c.Void, InTypes: _typs_P}
	_func_igGetWindowContentRegionMin_                           = &c.FuncPrototype{Name: "igGetWindowContentRegionMin", OutType: c.Void, InTypes: _typs_P}
	_func_igGetWindowContentRegionMax_                           = &c.FuncPrototype{Name: "igGetWindowContentRegionMax", OutType: c.Void, InTypes: _typs_P}
	_func_igSetScrollX_Float_                                    = &c.FuncPrototype{Name: "igSetScrollX_Float", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetScrollY_Float_                                    = &c.FuncPrototype{Name: "igSetScrollY_Float", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetScrollHereX_                                      = &c.FuncPrototype{Name: "igSetScrollHereX", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetScrollHereY_                                      = &c.FuncPrototype{Name: "igSetScrollHereY", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetScrollFromPosX_Float_                             = &c.FuncPrototype{Name: "igSetScrollFromPosX_Float", OutType: c.Void, InTypes: _typs_F32F32}
	_func_igSetScrollFromPosY_Float_                             = &c.FuncPrototype{Name: "igSetScrollFromPosY_Float", OutType: c.Void, InTypes: _typs_F32F32}
	_func_igPushFont_                                            = &c.FuncPrototype{Name: "igPushFont", OutType: c.Void, InTypes: _typs_P}
	_func_igPopFont_                                             = &c.FuncPrototype{Name: "igPopFont", OutType: c.Void, InTypes: nil}
	_func_igPushStyleColor_U32_                                  = &c.FuncPrototype{Name: "igPushStyleColor_U32", OutType: c.Void, InTypes: _typs_U32U32}
	_func_igPushStyleColor_Vec4_                                 = &c.FuncPrototype{Name: "igPushStyleColor_Vec4", OutType: c.Void, InTypes: _typs_U32Vec4}
	_func_igPopStyleColor_                                       = &c.FuncPrototype{Name: "igPopStyleColor", OutType: c.Void, InTypes: _typs_I32}
	_func_igPushStyleVar_Float_                                  = &c.FuncPrototype{Name: "igPushStyleVar_Float", OutType: c.Void, InTypes: _typs_I32F32}
	_func_igPushStyleVar_Vec2_                                   = &c.FuncPrototype{Name: "igPushStyleVar_Vec2", OutType: c.Void, InTypes: _typs_I32Vec2}
	_func_igPopStyleVar_                                         = &c.FuncPrototype{Name: "igPopStyleVar", OutType: c.Void, InTypes: _typs_I32}
	_func_igPushTabStop_                                         = &c.FuncPrototype{Name: "igPushTabStop", OutType: c.Void, InTypes: _typs_I32}
	_func_igPushButtonRepeat_                                    = &c.FuncPrototype{Name: "igPushButtonRepeat", OutType: c.Void, InTypes: _typs_I32}
	_func_igPushItemWidth_                                       = &c.FuncPrototype{Name: "igPushItemWidth", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetNextItemWidth_                                    = &c.FuncPrototype{Name: "igSetNextItemWidth", OutType: c.Void, InTypes: _typs_F32}
	_func_igPushTextWrapPos_                                     = &c.FuncPrototype{Name: "igPushTextWrapPos", OutType: c.Void, InTypes: _typs_F32}
	_func_igGetFontTexUvWhitePixel_                              = &c.FuncPrototype{Name: "igGetFontTexUvWhitePixel", OutType: c.Void, InTypes: _typs_P}
	_func_igSameLine_                                            = &c.FuncPrototype{Name: "igSameLine", OutType: c.Void, InTypes: _typs_F32F32}
	_func_igDummy_                                               = &c.FuncPrototype{Name: "igDummy", OutType: c.Void, InTypes: _typs_Vec2}
	_func_igIndent_                                              = &c.FuncPrototype{Name: "igIndent", OutType: c.Void, InTypes: _typs_F32}
	_func_igUnindent_                                            = &c.FuncPrototype{Name: "igUnindent", OutType: c.Void, InTypes: _typs_F32}
	_func_igGetCursorPos_                                        = &c.FuncPrototype{Name: "igGetCursorPos", OutType: c.Void, InTypes: _typs_P}
	_func_igSetCursorPos_                                        = &c.FuncPrototype{Name: "igSetCursorPos", OutType: c.Void, InTypes: _typs_Vec2}
	_func_igSetCursorPosX_                                       = &c.FuncPrototype{Name: "igSetCursorPosX", OutType: c.Void, InTypes: _typs_F32}
	_func_igSetCursorPosY_                                       = &c.FuncPrototype{Name: "igSetCursorPosY", OutType: c.Void, InTypes: _typs_F32}
	_func_igGetCursorStartPos_                                   = &c.FuncPrototype{Name: "igGetCursorStartPos", OutType: c.Void, InTypes: _typs_P}
	_func_igGetCursorScreenPos_                                  = &c.FuncPrototype{Name: "igGetCursorScreenPos", OutType: c.Void, InTypes: _typs_P}
	_func_igSetCursorScreenPos_                                  = &c.FuncPrototype{Name: "igSetCursorScreenPos", OutType: c.Void, InTypes: _typs_Vec2}
	_func_igAlignTextToFramePadding_                             = &c.FuncPrototype{Name: "igAlignTextToFramePadding", OutType: c.Void, InTypes: nil}
	_func_igPushID_Str_                                          = &c.FuncPrototype{Name: "igPushID_Str", OutType: c.Void, InTypes: _typs_P}
	_func_igPushID_StrStr_                                       = &c.FuncPrototype{Name: "igPushID_StrStr", OutType: c.Void, InTypes: _typs_PP}
	_func_igPushID_Ptr_                                          = &c.FuncPrototype{Name: "igPushID_Ptr", OutType: c.Void, InTypes: _typs_P}
	_func_igPushID_Int_                                          = &c.FuncPrototype{Name: "igPushID_Int", OutType: c.Void, InTypes: _typs_I32}
	_func_igTextUnformatted_                                     = &c.FuncPrototype{Name: "igTextUnformatted", OutType: c.Void, InTypes: _typs_PP}
	_func_igText_                                                = &c.FuncPrototype{Name: "igText", OutType: c.Void, InTypes: _typs_P}
	_func_igTextColored_                                         = &c.FuncPrototype{Name: "igTextColored", OutType: c.Void, InTypes: _typs_Vec4P}
	_func_igTextDisabled_                                        = &c.FuncPrototype{Name: "igTextDisabled", OutType: c.Void, InTypes: _typs_P}
	_func_igTextWrapped_                                         = &c.FuncPrototype{Name: "igTextWrapped", OutType: c.Void, InTypes: _typs_P}
	_func_igLabelText_                                           = &c.FuncPrototype{Name: "igLabelText", OutType: c.Void, InTypes: _typs_PP}
	_func_igBulletText_                                          = &c.FuncPrototype{Name: "igBulletText", OutType: c.Void, InTypes: _typs_P}
	_func_igSeparatorText_                                       = &c.FuncPrototype{Name: "igSeparatorText", OutType: c.Void, InTypes: _typs_P}
	_func_igProgressBar_                                         = &c.FuncPrototype{Name: "igProgressBar", OutType: c.Void, InTypes: _typs_F32Vec2P}
	_func_igImage_                                               = &c.FuncPrototype{Name: "igImage", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec4Vec4}
	_func_igSetColorEditOptions_                                 = &c.FuncPrototype{Name: "igSetColorEditOptions", OutType: c.Void, InTypes: _typs_U32}
	_func_igTreePush_Str_                                        = &c.FuncPrototype{Name: "igTreePush_Str", OutType: c.Void, InTypes: _typs_P}
	_func_igTreePush_Ptr_                                        = &c.FuncPrototype{Name: "igTreePush_Ptr", OutType: c.Void, InTypes: _typs_P}
	_func_igSetNextItemOpen_                                     = &c.FuncPrototype{Name: "igSetNextItemOpen", OutType: c.Void, InTypes: _typs_I32U32}
	_func_igPlotLines_FloatPtr_                                  = &c.FuncPrototype{Name: "igPlotLines_FloatPtr", OutType: c.Void, InTypes: _typs_PPI32I32PF32F32Vec2I32}
	_func_igPlotLines_FnFloatPtr_                                = &c.FuncPrototype{Name: "igPlotLines_FnFloatPtr", OutType: c.Void, InTypes: _typs_PPPI32I32PF32F32Vec2}
	_func_igPlotHistogram_FloatPtr_                              = &c.FuncPrototype{Name: "igPlotHistogram_FloatPtr", OutType: c.Void, InTypes: _typs_PPI32I32PF32F32Vec2I32}
	_func_igPlotHistogram_FnFloatPtr_                            = &c.FuncPrototype{Name: "igPlotHistogram_FnFloatPtr", OutType: c.Void, InTypes: _typs_PPPI32I32PF32F32Vec2}
	_func_igValue_Bool_                                          = &c.FuncPrototype{Name: "igValue_Bool", OutType: c.Void, InTypes: _typs_PI32}
	_func_igValue_Int_                                           = &c.FuncPrototype{Name: "igValue_Int", OutType: c.Void, InTypes: _typs_PI32}
	_func_igValue_Uint_                                          = &c.FuncPrototype{Name: "igValue_Uint", OutType: c.Void, InTypes: _typs_PU32}
	_func_igValue_Float_                                         = &c.FuncPrototype{Name: "igValue_Float", OutType: c.Void, InTypes: _typs_PF32P}
	_func_igEndMenuBar_                                          = &c.FuncPrototype{Name: "igEndMenuBar", OutType: c.Void, InTypes: nil}
	_func_igEndMainMenuBar_                                      = &c.FuncPrototype{Name: "igEndMainMenuBar", OutType: c.Void, InTypes: nil}
	_func_igEndMenu_                                             = &c.FuncPrototype{Name: "igEndMenu", OutType: c.Void, InTypes: nil}
	_func_igEndTooltip_                                          = &c.FuncPrototype{Name: "igEndTooltip", OutType: c.Void, InTypes: nil}
	_func_igSetTooltip_                                          = &c.FuncPrototype{Name: "igSetTooltip", OutType: c.Void, InTypes: _typs_P}
	_func_igSetItemTooltip_                                      = &c.FuncPrototype{Name: "igSetItemTooltip", OutType: c.Void, InTypes: _typs_P}
	_func_igEndPopup_                                            = &c.FuncPrototype{Name: "igEndPopup", OutType: c.Void, InTypes: nil}
	_func_igOpenPopup_Str_                                       = &c.FuncPrototype{Name: "igOpenPopup_Str", OutType: c.Void, InTypes: _typs_PU32}
	_func_igOpenPopup_ID_                                        = &c.FuncPrototype{Name: "igOpenPopup_ID", OutType: c.Void, InTypes: _typs_U32U32}
	_func_igOpenPopupOnItemClick_                                = &c.FuncPrototype{Name: "igOpenPopupOnItemClick", OutType: c.Void, InTypes: _typs_PU32}
	_func_igCloseCurrentPopup_                                   = &c.FuncPrototype{Name: "igCloseCurrentPopup", OutType: c.Void, InTypes: nil}
	_func_igTableNextRow_                                        = &c.FuncPrototype{Name: "igTableNextRow", OutType: c.Void, InTypes: _typs_U32F32}
	_func_igTableSetupColumn_                                    = &c.FuncPrototype{Name: "igTableSetupColumn", OutType: c.Void, InTypes: _typs_PU32F32U32}
	_func_igTableSetupScrollFreeze_                              = &c.FuncPrototype{Name: "igTableSetupScrollFreeze", OutType: c.Void, InTypes: _typs_I32I32}
	_func_igTableHeader_                                         = &c.FuncPrototype{Name: "igTableHeader", OutType: c.Void, InTypes: _typs_P}
	_func_igTableSetColumnEnabled_                               = &c.FuncPrototype{Name: "igTableSetColumnEnabled", OutType: c.Void, InTypes: _typs_I32I32}
	_func_igTableSetBgColor_                                     = &c.FuncPrototype{Name: "igTableSetBgColor", OutType: c.Void, InTypes: _typs_U32U32I32}
	_func_igColumns_                                             = &c.FuncPrototype{Name: "igColumns", OutType: c.Void, InTypes: _typs_I32PI32}
	_func_igSetColumnWidth_                                      = &c.FuncPrototype{Name: "igSetColumnWidth", OutType: c.Void, InTypes: _typs_I32F32}
	_func_igSetColumnOffset_                                     = &c.FuncPrototype{Name: "igSetColumnOffset", OutType: c.Void, InTypes: _typs_I32F32}
	_func_igSetTabItemClosed_                                    = &c.FuncPrototype{Name: "igSetTabItemClosed", OutType: c.Void, InTypes: _typs_P}
	_func_igSetNextWindowDockID_                                 = &c.FuncPrototype{Name: "igSetNextWindowDockID", OutType: c.Void, InTypes: _typs_U32U32}
	_func_igSetNextWindowClass_                                  = &c.FuncPrototype{Name: "igSetNextWindowClass", OutType: c.Void, InTypes: _typs_P}
	_func_igLogToTTY_                                            = &c.FuncPrototype{Name: "igLogToTTY", OutType: c.Void, InTypes: _typs_I32}
	_func_igLogToFile_                                           = &c.FuncPrototype{Name: "igLogToFile", OutType: c.Void, InTypes: _typs_I32P}
	_func_igLogToClipboard_                                      = &c.FuncPrototype{Name: "igLogToClipboard", OutType: c.Void, InTypes: _typs_I32}
	_func_igLogTextV_                                            = &c.FuncPrototype{Name: "igLogTextV", OutType: c.Void, InTypes: _typs_P}
	_func_igBeginDisabled_                                       = &c.FuncPrototype{Name: "igBeginDisabled", OutType: c.Void, InTypes: _typs_I32}
	_func_igPushClipRect_                                        = &c.FuncPrototype{Name: "igPushClipRect", OutType: c.Void, InTypes: _typs_Vec2Vec2I32}
	_func_igSetKeyboardFocusHere_                                = &c.FuncPrototype{Name: "igSetKeyboardFocusHere", OutType: c.Void, InTypes: _typs_I32}
	_func_igGetItemRectMin_                                      = &c.FuncPrototype{Name: "igGetItemRectMin", OutType: c.Void, InTypes: _typs_P}
	_func_igGetItemRectMax_                                      = &c.FuncPrototype{Name: "igGetItemRectMax", OutType: c.Void, InTypes: _typs_P}
	_func_igGetItemRectSize_                                     = &c.FuncPrototype{Name: "igGetItemRectSize", OutType: c.Void, InTypes: _typs_P}
	_func_igSetStateStorage_                                     = &c.FuncPrototype{Name: "igSetStateStorage", OutType: c.Void, InTypes: _typs_P}
	_func_igCalcTextSize_                                        = &c.FuncPrototype{Name: "igCalcTextSize", OutType: c.Void, InTypes: _typs_PPPI32F32}
	_func_igColorConvertU32ToFloat4_                             = &c.FuncPrototype{Name: "igColorConvertU32ToFloat4", OutType: c.Void, InTypes: _typs_PU32}
	_func_igColorConvertRGBtoHSV_                                = &c.FuncPrototype{Name: "igColorConvertRGBtoHSV", OutType: c.Void, InTypes: _typs_F32F32F32PPP}
	_func_igColorConvertHSVtoRGB_                                = &c.FuncPrototype{Name: "igColorConvertHSVtoRGB", OutType: c.Void, InTypes: _typs_F32F32F32PPP}
	_func_igSetNextFrameWantCaptureKeyboard_                     = &c.FuncPrototype{Name: "igSetNextFrameWantCaptureKeyboard", OutType: c.Void, InTypes: _typs_I32}
	_func_igGetMousePos_                                         = &c.FuncPrototype{Name: "igGetMousePos", OutType: c.Void, InTypes: _typs_P}
	_func_igGetMousePosOnOpeningCurrentPopup_                    = &c.FuncPrototype{Name: "igGetMousePosOnOpeningCurrentPopup", OutType: c.Void, InTypes: _typs_P}
	_func_igGetMouseDragDelta_                                   = &c.FuncPrototype{Name: "igGetMouseDragDelta", OutType: c.Void, InTypes: _typs_PU32F32}
	_func_igResetMouseDragDelta_                                 = &c.FuncPrototype{Name: "igResetMouseDragDelta", OutType: c.Void, InTypes: _typs_U32}
	_func_igSetMouseCursor_                                      = &c.FuncPrototype{Name: "igSetMouseCursor", OutType: c.Void, InTypes: _typs_I32}
	_func_igSetNextFrameWantCaptureMouse_                        = &c.FuncPrototype{Name: "igSetNextFrameWantCaptureMouse", OutType: c.Void, InTypes: _typs_I32}
	_func_igSetClipboardText_                                    = &c.FuncPrototype{Name: "igSetClipboardText", OutType: c.Void, InTypes: _typs_P}
	_func_igLoadIniSettingsFromDisk_                             = &c.FuncPrototype{Name: "igLoadIniSettingsFromDisk", OutType: c.Void, InTypes: _typs_P}
	_func_igLoadIniSettingsFromMemory_                           = &c.FuncPrototype{Name: "igLoadIniSettingsFromMemory", OutType: c.Void, InTypes: _typs_PU64}
	_func_igSaveIniSettingsToDisk_                               = &c.FuncPrototype{Name: "igSaveIniSettingsToDisk", OutType: c.Void, InTypes: _typs_P}
	_func_igDebugTextEncoding_                                   = &c.FuncPrototype{Name: "igDebugTextEncoding", OutType: c.Void, InTypes: _typs_P}
	_func_igRenderPlatformWindowsDefault_                        = &c.FuncPrototype{Name: "igRenderPlatformWindowsDefault", OutType: c.Void, InTypes: _typs_PP}
	_func_ImGuiStyle_destroy_                                    = &c.FuncPrototype{Name: "ImGuiStyle_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiStyle_ScaleAllSizes_                              = &c.FuncPrototype{Name: "ImGuiStyle_ScaleAllSizes", OutType: c.Void, InTypes: _typs_PF32}
	_func_ImGuiIO_AddKeyEvent_                                   = &c.FuncPrototype{Name: "ImGuiIO_AddKeyEvent", OutType: c.Void, InTypes: _typs_PU32I32}
	_func_ImGuiIO_AddKeyAnalogEvent_                             = &c.FuncPrototype{Name: "ImGuiIO_AddKeyAnalogEvent", OutType: c.Void, InTypes: _typs_PU32I32F32}
	_func_ImGuiIO_AddMousePosEvent_                              = &c.FuncPrototype{Name: "ImGuiIO_AddMousePosEvent", OutType: c.Void, InTypes: _typs_PF32F32}
	_func_ImGuiIO_AddMouseButtonEvent_                           = &c.FuncPrototype{Name: "ImGuiIO_AddMouseButtonEvent", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_ImGuiIO_AddMouseWheelEvent_                            = &c.FuncPrototype{Name: "ImGuiIO_AddMouseWheelEvent", OutType: c.Void, InTypes: _typs_PF32F32}
	_func_ImGuiIO_AddMouseSourceEvent_                           = &c.FuncPrototype{Name: "ImGuiIO_AddMouseSourceEvent", OutType: c.Void, InTypes: _typs_PU32}
	_func_ImGuiIO_AddMouseViewportEvent_                         = &c.FuncPrototype{Name: "ImGuiIO_AddMouseViewportEvent", OutType: c.Void, InTypes: _typs_PU32}
	_func_ImGuiIO_AddFocusEvent_                                 = &c.FuncPrototype{Name: "ImGuiIO_AddFocusEvent", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImGuiIO_AddInputCharacter_                             = &c.FuncPrototype{Name: "ImGuiIO_AddInputCharacter", OutType: c.Void, InTypes: _typs_PU32}
	_func_ImGuiIO_AddInputCharacterUTF16_                        = &c.FuncPrototype{Name: "ImGuiIO_AddInputCharacterUTF16", OutType: c.Void, InTypes: _typs_PU16}
	_func_ImGuiIO_AddInputCharactersUTF8_                        = &c.FuncPrototype{Name: "ImGuiIO_AddInputCharactersUTF8", OutType: c.Void, InTypes: _typs_PP}
	_func_ImGuiIO_SetKeyEventNativeData_                         = &c.FuncPrototype{Name: "ImGuiIO_SetKeyEventNativeData", OutType: c.Void, InTypes: _typs_PU32I32I32I32}
	_func_ImGuiIO_SetAppAcceptingEvents_                         = &c.FuncPrototype{Name: "ImGuiIO_SetAppAcceptingEvents", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImGuiIO_ClearInputCharacters_                          = &c.FuncPrototype{Name: "ImGuiIO_ClearInputCharacters", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiIO_ClearInputKeys_                                = &c.FuncPrototype{Name: "ImGuiIO_ClearInputKeys", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiIO_destroy_                                       = &c.FuncPrototype{Name: "ImGuiIO_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiInputTextCallbackData_destroy_                    = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiInputTextCallbackData_DeleteChars_                = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_DeleteChars", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_ImGuiInputTextCallbackData_InsertChars_                = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_InsertChars", OutType: c.Void, InTypes: _typs_PI32PP}
	_func_ImGuiInputTextCallbackData_SelectAll_                  = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_SelectAll", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiInputTextCallbackData_ClearSelection_             = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_ClearSelection", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiInputTextCallbackData_HasSelection_               = &c.FuncPrototype{Name: "ImGuiInputTextCallbackData_HasSelection", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiWindowClass_ImGuiWindowClass_                     = &c.FuncPrototype{Name: "ImGuiWindowClass_ImGuiWindowClass", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiWindowClass_destroy_                              = &c.FuncPrototype{Name: "ImGuiWindowClass_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontConfig_destroy_                                  = &c.FuncPrototype{Name: "ImFontConfig_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontAtlas_destroy_                                   = &c.FuncPrototype{Name: "ImFontAtlas_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiPayload_ImGuiPayload_                             = &c.FuncPrototype{Name: "ImGuiPayload_ImGuiPayload", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiPayload_destroy_                                  = &c.FuncPrototype{Name: "ImGuiPayload_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiPayload_Clear_                                    = &c.FuncPrototype{Name: "ImGuiPayload_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiPayload_IsDataType_                               = &c.FuncPrototype{Name: "ImGuiPayload_IsDataType", OutType: c.I32, InTypes: _typs_PP}
	_func_ImGuiPayload_IsPreview_                                = &c.FuncPrototype{Name: "ImGuiPayload_IsPreview", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiPayload_IsDelivery_                               = &c.FuncPrototype{Name: "ImGuiPayload_IsDelivery", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs_   = &c.FuncPrototype{Name: "ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiTableColumnSortSpecs_destroy_                     = &c.FuncPrototype{Name: "ImGuiTableColumnSortSpecs_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTableSortSpecs_ImGuiTableSortSpecs_               = &c.FuncPrototype{Name: "ImGuiTableSortSpecs_ImGuiTableSortSpecs", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiTableSortSpecs_destroy_                           = &c.FuncPrototype{Name: "ImGuiTableSortSpecs_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiOnceUponAFrame_ImGuiOnceUponAFrame_               = &c.FuncPrototype{Name: "ImGuiOnceUponAFrame_ImGuiOnceUponAFrame", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiOnceUponAFrame_destroy_                           = &c.FuncPrototype{Name: "ImGuiOnceUponAFrame_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextFilter_ImGuiTextFilter_                       = &c.FuncPrototype{Name: "ImGuiTextFilter_ImGuiTextFilter", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImGuiTextFilter_destroy_                               = &c.FuncPrototype{Name: "ImGuiTextFilter_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextFilter_Draw_                                  = &c.FuncPrototype{Name: "ImGuiTextFilter_Draw", OutType: c.I32, InTypes: _typs_PPF32}
	_func_ImGuiTextFilter_PassFilter_                            = &c.FuncPrototype{Name: "ImGuiTextFilter_PassFilter", OutType: c.I32, InTypes: _typs_PPP}
	_func_ImGuiTextFilter_Build_                                 = &c.FuncPrototype{Name: "ImGuiTextFilter_Build", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextFilter_Clear_                                 = &c.FuncPrototype{Name: "ImGuiTextFilter_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextFilter_IsActive_                              = &c.FuncPrototype{Name: "ImGuiTextFilter_IsActive", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiTextRange_ImGuiTextRange_Nil_                     = &c.FuncPrototype{Name: "ImGuiTextRange_ImGuiTextRange_Nil", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiTextRange_destroy_                                = &c.FuncPrototype{Name: "ImGuiTextRange_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextRange_ImGuiTextRange_Str_                     = &c.FuncPrototype{Name: "ImGuiTextRange_ImGuiTextRange_Str", OutType: c.Pointer, InTypes: _typs_PP}
	_func_ImGuiTextRange_empty_                                  = &c.FuncPrototype{Name: "ImGuiTextRange_empty", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiTextRange_split_                                  = &c.FuncPrototype{Name: "ImGuiTextRange_split", OutType: c.Void, InTypes: _typs_PU8P}
	_func_ImGuiTextBuffer_ImGuiTextBuffer_                       = &c.FuncPrototype{Name: "ImGuiTextBuffer_ImGuiTextBuffer", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiTextBuffer_destroy_                               = &c.FuncPrototype{Name: "ImGuiTextBuffer_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextBuffer_begin_                                 = &c.FuncPrototype{Name: "ImGuiTextBuffer_begin", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImColor_destroy_                                       = &c.FuncPrototype{Name: "ImColor_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImColor_HSV_                                           = &c.FuncPrototype{Name: "ImColor_HSV", OutType: c.Void, InTypes: _typs_PF32F32F32F32}
	_func_ImColor_ImColor_Float_                                 = &c.FuncPrototype{Name: "ImColor_ImColor_Float", OutType: c.Pointer, InTypes: _typs_F32F32F32F32}
	_func_ImColor_ImColor_Int_                                   = &c.FuncPrototype{Name: "ImColor_ImColor_Int", OutType: c.Pointer, InTypes: _typs_I32I32I32I32}
	_func_ImColor_ImColor_Nil_                                   = &c.FuncPrototype{Name: "ImColor_ImColor_Nil", OutType: c.Pointer, InTypes: nil}
	_func_ImColor_ImColor_U32_                                   = &c.FuncPrototype{Name: "ImColor_ImColor_U32", OutType: c.Pointer, InTypes: _typs_U32}
	_func_ImColor_ImColor_Vec4_                                  = &c.FuncPrototype{Name: "ImColor_ImColor_Vec4", OutType: c.Pointer, InTypes: _typs_Vec4}
	_func_ImColor_SetHSV_                                        = &c.FuncPrototype{Name: "ImColor_SetHSV", OutType: c.Void, InTypes: _typs_PF32F32F32F32}
	_func_ImDrawCmd_destroy_                                     = &c.FuncPrototype{Name: "ImDrawCmd_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawCmd_GetTexID_                                    = &c.FuncPrototype{Name: "ImDrawCmd_GetTexID", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImDrawCmd_ImDrawCmd_                                   = &c.FuncPrototype{Name: "ImDrawCmd_ImDrawCmd", OutType: c.Pointer, InTypes: nil}
	_func_ImDrawData_Clear_                                      = &c.FuncPrototype{Name: "ImDrawData_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawData_DeIndexAllBuffers_                          = &c.FuncPrototype{Name: "ImDrawData_DeIndexAllBuffers", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawData_destroy_                                    = &c.FuncPrototype{Name: "ImDrawData_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawData_ImDrawData_                                 = &c.FuncPrototype{Name: "ImDrawData_ImDrawData", OutType: c.Pointer, InTypes: nil}
	_func_ImDrawData_ScaleClipRects_                             = &c.FuncPrototype{Name: "ImDrawData_ScaleClipRects", OutType: c.Void, InTypes: _typs_PVec2}
	_func_ImDrawList__CalcCircleAutoSegmentCount_                = &c.FuncPrototype{Name: "ImDrawList__CalcCircleAutoSegmentCount", OutType: c.I32, InTypes: _typs_PF32}
	_func_ImDrawList__ClearFreeMemory_                           = &c.FuncPrototype{Name: "ImDrawList__ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList__OnChangedClipRect_                         = &c.FuncPrototype{Name: "ImDrawList__OnChangedClipRect", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList__OnChangedTextureID_                        = &c.FuncPrototype{Name: "ImDrawList__OnChangedTextureID", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList__OnChangedVtxOffset_                        = &c.FuncPrototype{Name: "ImDrawList__OnChangedVtxOffset", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList__PathArcToFastEx_                           = &c.FuncPrototype{Name: "ImDrawList__PathArcToFastEx", OutType: c.Void, InTypes: _typs_PVec2F32I32I32I32}
	_func_ImDrawList__PathArcToN_                                = &c.FuncPrototype{Name: "ImDrawList__PathArcToN", OutType: c.Void, InTypes: _typs_PVec2F32F32F32I32}
	_func_ImDrawList__PopUnusedDrawCmd_                          = &c.FuncPrototype{Name: "ImDrawList__PopUnusedDrawCmd", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList__ResetForNewFrame_                          = &c.FuncPrototype{Name: "ImDrawList__ResetForNewFrame", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList__TryMergeDrawCmds_                          = &c.FuncPrototype{Name: "ImDrawList__TryMergeDrawCmds", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_AddBezierCubic_                             = &c.FuncPrototype{Name: "ImDrawList_AddBezierCubic", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32F32I32}
	_func_ImDrawList_AddBezierQuadratic_                         = &c.FuncPrototype{Name: "ImDrawList_AddBezierQuadratic", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2U32F32I32}
	_func_ImDrawList_AddCircle_                                  = &c.FuncPrototype{Name: "ImDrawList_AddCircle", OutType: c.Void, InTypes: _typs_PVec2F32U32I32F32}
	_func_ImDrawList_AddCircleFilled_                            = &c.FuncPrototype{Name: "ImDrawList_AddCircleFilled", OutType: c.Void, InTypes: _typs_PVec2F32U32I32}
	_func_ImDrawList_AddConvexPolyFilled_                        = &c.FuncPrototype{Name: "ImDrawList_AddConvexPolyFilled", OutType: c.Void, InTypes: _typs_PPI32U32}
	_func_ImDrawList_AddDrawCmd_                                 = &c.FuncPrototype{Name: "ImDrawList_AddDrawCmd", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_AddImage_                                   = &c.FuncPrototype{Name: "ImDrawList_AddImage", OutType: c.Void, InTypes: _typs_PPVec2Vec2Vec2Vec2U32}
	_func_ImDrawList_AddImageQuad_                               = &c.FuncPrototype{Name: "ImDrawList_AddImageQuad", OutType: c.Void, InTypes: _typs_PPVec2Vec2Vec2Vec2Vec2Vec2Vec2Vec2U32}
	_func_ImDrawList_AddImageRounded_                            = &c.FuncPrototype{Name: "ImDrawList_AddImageRounded", OutType: c.Void, InTypes: _typs_PPVec2Vec2Vec2Vec2U32F32U32}
	_func_ImDrawList_AddLine_                                    = &c.FuncPrototype{Name: "ImDrawList_AddLine", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32}
	_func_ImDrawList_AddNgon_                                    = &c.FuncPrototype{Name: "ImDrawList_AddNgon", OutType: c.Void, InTypes: _typs_PVec2F32U32I32F32}
	_func_ImDrawList_AddNgonFilled_                              = &c.FuncPrototype{Name: "ImDrawList_AddNgonFilled", OutType: c.Void, InTypes: _typs_PVec2F32U32I32}
	_func_ImDrawList_AddPolyline_                                = &c.FuncPrototype{Name: "ImDrawList_AddPolyline", OutType: c.Void, InTypes: _typs_PPI32U32U32F32}
	_func_ImDrawList_AddQuad_                                    = &c.FuncPrototype{Name: "ImDrawList_AddQuad", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32F32}
	_func_ImDrawList_AddQuadFilled_                              = &c.FuncPrototype{Name: "ImDrawList_AddQuadFilled", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32}
	_func_ImDrawList_AddRect_                                    = &c.FuncPrototype{Name: "ImDrawList_AddRect", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32U32F32}
	_func_ImDrawList_AddRectFilled_                              = &c.FuncPrototype{Name: "ImDrawList_AddRectFilled", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32U32}
	_func_ImDrawList_AddRectFilledMultiColor_                    = &c.FuncPrototype{Name: "ImDrawList_AddRectFilledMultiColor", OutType: c.Void, InTypes: _typs_PVec2Vec2U32U32U32U32}
	_func_ImDrawList_AddText_FontPtr_                            = &c.FuncPrototype{Name: "ImDrawList_AddText_FontPtr", OutType: c.Void, InTypes: _typs_PPF32Vec2U32PPF32P}
	_func_ImDrawList_AddText_Vec2_                               = &c.FuncPrototype{Name: "ImDrawList_AddText_Vec2", OutType: c.Void, InTypes: _typs_PVec2U32PP}
	_func_ImDrawList_AddTriangle_                                = &c.FuncPrototype{Name: "ImDrawList_AddTriangle", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2U32F32}
	_func_ImDrawList_AddTriangleFilled_                          = &c.FuncPrototype{Name: "ImDrawList_AddTriangleFilled", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2U32}
	_func_ImDrawList_ChannelsMerge_                              = &c.FuncPrototype{Name: "ImDrawList_ChannelsMerge", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_ChannelsSetCurrent_                         = &c.FuncPrototype{Name: "ImDrawList_ChannelsSetCurrent", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImDrawList_ChannelsSplit_                              = &c.FuncPrototype{Name: "ImDrawList_ChannelsSplit", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImDrawList_CloneOutput_                                = &c.FuncPrototype{Name: "ImDrawList_CloneOutput", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImDrawList_destroy_                                    = &c.FuncPrototype{Name: "ImDrawList_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_GetClipRectMax_                             = &c.FuncPrototype{Name: "ImDrawList_GetClipRectMax", OutType: c.Void, InTypes: _typs_PP}
	_func_ImDrawList_GetClipRectMin_                             = &c.FuncPrototype{Name: "ImDrawList_GetClipRectMin", OutType: c.Void, InTypes: _typs_PP}
	_func_ImDrawList_ImDrawList_                                 = &c.FuncPrototype{Name: "ImDrawList_ImDrawList", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImDrawList_PathArcTo_                                  = &c.FuncPrototype{Name: "ImDrawList_PathArcTo", OutType: c.Void, InTypes: _typs_PVec2F32F32F32I32}
	_func_ImDrawList_PathArcToFast_                              = &c.FuncPrototype{Name: "ImDrawList_PathArcToFast", OutType: c.Void, InTypes: _typs_PVec2F32I32I32}
	_func_ImDrawList_PathBezierCubicCurveTo_                     = &c.FuncPrototype{Name: "ImDrawList_PathBezierCubicCurveTo", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2I32}
	_func_ImDrawList_PathBezierQuadraticCurveTo_                 = &c.FuncPrototype{Name: "ImDrawList_PathBezierQuadraticCurveTo", OutType: c.Void, InTypes: _typs_PVec2Vec2I32}
	_func_ImDrawList_PathClear_                                  = &c.FuncPrototype{Name: "ImDrawList_PathClear", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_PathFillConvex_                             = &c.FuncPrototype{Name: "ImDrawList_PathFillConvex", OutType: c.Void, InTypes: _typs_PU32}
	_func_ImDrawList_PathLineTo_                                 = &c.FuncPrototype{Name: "ImDrawList_PathLineTo", OutType: c.Void, InTypes: _typs_PVec2}
	_func_ImDrawList_PathLineToMergeDuplicate_                   = &c.FuncPrototype{Name: "ImDrawList_PathLineToMergeDuplicate", OutType: c.Void, InTypes: _typs_PVec2}
	_func_ImDrawList_PathRect_                                   = &c.FuncPrototype{Name: "ImDrawList_PathRect", OutType: c.Void, InTypes: _typs_PVec2Vec2F32U32}
	_func_ImDrawList_PathStroke_                                 = &c.FuncPrototype{Name: "ImDrawList_PathStroke", OutType: c.Void, InTypes: _typs_PU32U32F32}
	_func_ImDrawList_PopClipRect_                                = &c.FuncPrototype{Name: "ImDrawList_PopClipRect", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_PopTextureID_                               = &c.FuncPrototype{Name: "ImDrawList_PopTextureID", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_PrimQuadUV_                                 = &c.FuncPrototype{Name: "ImDrawList_PrimQuadUV", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2Vec2Vec2Vec2Vec2U32}
	_func_ImDrawList_PrimRect_                                   = &c.FuncPrototype{Name: "ImDrawList_PrimRect", OutType: c.Void, InTypes: _typs_PVec2Vec2U32}
	_func_ImDrawList_PrimRectUV_                                 = &c.FuncPrototype{Name: "ImDrawList_PrimRectUV", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32}
	_func_ImDrawList_PrimReserve_                                = &c.FuncPrototype{Name: "ImDrawList_PrimReserve", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_ImDrawList_PrimUnreserve_                              = &c.FuncPrototype{Name: "ImDrawList_PrimUnreserve", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_ImDrawList_PrimVtx_                                    = &c.FuncPrototype{Name: "ImDrawList_PrimVtx", OutType: c.Void, InTypes: _typs_PVec2Vec2U32}
	_func_ImDrawList_PrimWriteIdx_                               = &c.FuncPrototype{Name: "ImDrawList_PrimWriteIdx", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_PrimWriteVtx_                               = &c.FuncPrototype{Name: "ImDrawList_PrimWriteVtx", OutType: c.Void, InTypes: _typs_PVec2Vec2U32}
	_func_ImDrawList_PushClipRect_                               = &c.FuncPrototype{Name: "ImDrawList_PushClipRect", OutType: c.Void, InTypes: _typs_PVec2Vec2I32}
	_func_ImDrawList_PushClipRectFullScreen_                     = &c.FuncPrototype{Name: "ImDrawList_PushClipRectFullScreen", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawList_PushTextureID_                              = &c.FuncPrototype{Name: "ImDrawList_PushTextureID", OutType: c.Void, InTypes: _typs_PP}
	_func_ImDrawListSplitter_Clear_                              = &c.FuncPrototype{Name: "ImDrawListSplitter_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawListSplitter_ClearFreeMemory_                    = &c.FuncPrototype{Name: "ImDrawListSplitter_ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawListSplitter_destroy_                            = &c.FuncPrototype{Name: "ImDrawListSplitter_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImDrawListSplitter_ImDrawListSplitter_                 = &c.FuncPrototype{Name: "ImDrawListSplitter_ImDrawListSplitter", OutType: c.Pointer, InTypes: nil}
	_func_ImDrawListSplitter_Merge_                              = &c.FuncPrototype{Name: "ImDrawListSplitter_Merge", OutType: c.Void, InTypes: _typs_PP}
	_func_ImDrawListSplitter_SetCurrentChannel_                  = &c.FuncPrototype{Name: "ImDrawListSplitter_SetCurrentChannel", OutType: c.Void, InTypes: _typs_PPI32}
	_func_ImDrawListSplitter_Split_                              = &c.FuncPrototype{Name: "ImDrawListSplitter_Split", OutType: c.Void, InTypes: _typs_PPI32}
	_func_ImFont_AddGlyph_                                       = &c.FuncPrototype{Name: "ImFont_AddGlyph", OutType: c.Void, InTypes: _typs_PPU16F32F32F32F32F32F32F32F32F32}
	_func_ImFont_AddRemapChar_                                   = &c.FuncPrototype{Name: "ImFont_AddRemapChar", OutType: c.Void, InTypes: _typs_PU16U16I32}
	_func_ImFont_BuildLookupTable_                               = &c.FuncPrototype{Name: "ImFont_BuildLookupTable", OutType: c.Void, InTypes: _typs_P}
	_func_ImFont_CalcWordWrapPositionA_                          = &c.FuncPrototype{Name: "ImFont_CalcWordWrapPositionA", OutType: c.Pointer, InTypes: _typs_PF32PPF32}
	_func_ImFont_ClearOutputData_                                = &c.FuncPrototype{Name: "ImFont_ClearOutputData", OutType: c.Void, InTypes: _typs_P}
	_func_ImFont_destroy_                                        = &c.FuncPrototype{Name: "ImFont_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImFont_FindGlyph_                                      = &c.FuncPrototype{Name: "ImFont_FindGlyph", OutType: c.Pointer, InTypes: _typs_PU16}
	_func_ImFont_FindGlyphNoFallback_                            = &c.FuncPrototype{Name: "ImFont_FindGlyphNoFallback", OutType: c.Pointer, InTypes: _typs_PU16}
	_func_ImFont_GetCharAdvance_                                 = &c.FuncPrototype{Name: "ImFont_GetCharAdvance", OutType: c.F32, InTypes: _typs_PU16}
	_func_ImFont_GetDebugName_                                   = &c.FuncPrototype{Name: "ImFont_GetDebugName", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImFont_GrowIndex_                                      = &c.FuncPrototype{Name: "ImFont_GrowIndex", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImFont_ImFont_                                         = &c.FuncPrototype{Name: "ImFont_ImFont", OutType: c.Pointer, InTypes: nil}
	_func_ImFont_IsGlyphRangeUnused_                             = &c.FuncPrototype{Name: "ImFont_IsGlyphRangeUnused", OutType: c.I32, InTypes: _typs_PU32U32}
	_func_ImFont_IsLoaded_                                       = &c.FuncPrototype{Name: "ImFont_IsLoaded", OutType: c.I32, InTypes: _typs_P}
	_func_ImFont_RenderChar_                                     = &c.FuncPrototype{Name: "ImFont_RenderChar", OutType: c.Void, InTypes: _typs_PPF32Vec2U32U16}
	_func_ImFont_RenderText_                                     = &c.FuncPrototype{Name: "ImFont_RenderText", OutType: c.Void, InTypes: _typs_PPF32Vec2U32Vec4PPF32I32}
	_func_ImFont_SetGlyphVisible_                                = &c.FuncPrototype{Name: "ImFont_SetGlyphVisible", OutType: c.Void, InTypes: _typs_PU16I32}
	_func_ImFontAtlas_AddCustomRectFontGlyph_                    = &c.FuncPrototype{Name: "ImFontAtlas_AddCustomRectFontGlyph", OutType: c.I32, InTypes: _typs_PPU16I32I32F32Vec2}
	_func_ImFontAtlas_AddCustomRectRegular_                      = &c.FuncPrototype{Name: "ImFontAtlas_AddCustomRectRegular", OutType: c.I32, InTypes: _typs_PI32I32}
	_func_ImFontAtlas_Build_                                     = &c.FuncPrototype{Name: "ImFontAtlas_Build", OutType: c.I32, InTypes: _typs_P}
	_func_ImFontAtlas_CalcCustomRectUV_                          = &c.FuncPrototype{Name: "ImFontAtlas_CalcCustomRectUV", OutType: c.Void, InTypes: _typs_PPPP}
	_func_ImFontAtlas_Clear_                                     = &c.FuncPrototype{Name: "ImFontAtlas_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontAtlas_ClearFonts_                                = &c.FuncPrototype{Name: "ImFontAtlas_ClearFonts", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontAtlas_ClearInputData_                            = &c.FuncPrototype{Name: "ImFontAtlas_ClearInputData", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontAtlas_ClearTexData_                              = &c.FuncPrototype{Name: "ImFontAtlas_ClearTexData", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontAtlas_GetCustomRectByIndex_                      = &c.FuncPrototype{Name: "ImFontAtlas_GetCustomRectByIndex", OutType: c.Pointer, InTypes: _typs_PI32}
	_func_ImFontAtlas_IsBuilt_                                   = &c.FuncPrototype{Name: "ImFontAtlas_IsBuilt", OutType: c.I32, InTypes: _typs_P}
	_func_ImFontAtlas_SetTexID_                                  = &c.FuncPrototype{Name: "ImFontAtlas_SetTexID", OutType: c.Void, InTypes: _typs_PP}
	_func_ImFontAtlasCustomRect_destroy_                         = &c.FuncPrototype{Name: "ImFontAtlasCustomRect_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontAtlasCustomRect_ImFontAtlasCustomRect_           = &c.FuncPrototype{Name: "ImFontAtlasCustomRect_ImFontAtlasCustomRect", OutType: c.Pointer, InTypes: nil}
	_func_ImFontAtlasCustomRect_IsPacked_                        = &c.FuncPrototype{Name: "ImFontAtlasCustomRect_IsPacked", OutType: c.I32, InTypes: _typs_P}
	_func_ImFontGlyphRangesBuilder_AddChar_                      = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_AddChar", OutType: c.Void, InTypes: _typs_PU16}
	_func_ImFontGlyphRangesBuilder_AddRanges_                    = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_AddRanges", OutType: c.Void, InTypes: _typs_PP}
	_func_ImFontGlyphRangesBuilder_AddText_                      = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_AddText", OutType: c.Void, InTypes: _typs_PPP}
	_func_ImFontGlyphRangesBuilder_BuildRanges_                  = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_BuildRanges", OutType: c.Void, InTypes: _typs_PP}
	_func_ImFontGlyphRangesBuilder_Clear_                        = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontGlyphRangesBuilder_destroy_                      = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImFontGlyphRangesBuilder_GetBit_                       = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_GetBit", OutType: c.I32, InTypes: _typs_PU64}
	_func_ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder_     = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder", OutType: c.Pointer, InTypes: nil}
	_func_ImFontGlyphRangesBuilder_SetBit_                       = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_SetBit", OutType: c.Void, InTypes: _typs_PU64}
	_func_ImGuiListClipper_Begin_                                = &c.FuncPrototype{Name: "ImGuiListClipper_Begin", OutType: c.Void, InTypes: _typs_PI32F32}
	_func_ImGuiListClipper_destroy_                              = &c.FuncPrototype{Name: "ImGuiListClipper_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiListClipper_End_                                  = &c.FuncPrototype{Name: "ImGuiListClipper_End", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiListClipper_ImGuiListClipper_                     = &c.FuncPrototype{Name: "ImGuiListClipper_ImGuiListClipper", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiListClipper_IncludeRangeByIndices_                = &c.FuncPrototype{Name: "ImGuiListClipper_IncludeRangeByIndices", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_ImGuiListClipper_Step_                                 = &c.FuncPrototype{Name: "ImGuiListClipper_Step", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiPlatformImeData_destroy_                          = &c.FuncPrototype{Name: "ImGuiPlatformImeData_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiPlatformImeData_ImGuiPlatformImeData_             = &c.FuncPrototype{Name: "ImGuiPlatformImeData_ImGuiPlatformImeData", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiPlatformIO_destroy_                               = &c.FuncPrototype{Name: "ImGuiPlatformIO_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiPlatformIO_ImGuiPlatformIO_                       = &c.FuncPrototype{Name: "ImGuiPlatformIO_ImGuiPlatformIO", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiPlatformMonitor_destroy_                          = &c.FuncPrototype{Name: "ImGuiPlatformMonitor_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiPlatformMonitor_ImGuiPlatformMonitor_             = &c.FuncPrototype{Name: "ImGuiPlatformMonitor_ImGuiPlatformMonitor", OutType: c.Pointer, InTypes: nil}
	_func_ImGuiStorage_BuildSortByKey_                           = &c.FuncPrototype{Name: "ImGuiStorage_BuildSortByKey", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiStorage_Clear_                                    = &c.FuncPrototype{Name: "ImGuiStorage_Clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiStorage_GetBool_                                  = &c.FuncPrototype{Name: "ImGuiStorage_GetBool", OutType: c.I32, InTypes: _typs_PU32I32}
	_func_ImGuiStorage_GetFloat_                                 = &c.FuncPrototype{Name: "ImGuiStorage_GetFloat", OutType: c.F32, InTypes: _typs_PU32F32}
	_func_ImGuiStorage_GetFloatRef_                              = &c.FuncPrototype{Name: "ImGuiStorage_GetFloatRef", OutType: c.Pointer, InTypes: _typs_PU32F32}
	_func_ImGuiStorage_GetInt_                                   = &c.FuncPrototype{Name: "ImGuiStorage_GetInt", OutType: c.I32, InTypes: _typs_PU32I32}
	_func_ImGuiStorage_GetIntRef_                                = &c.FuncPrototype{Name: "ImGuiStorage_GetIntRef", OutType: c.Pointer, InTypes: _typs_PU32I32}
	_func_ImGuiStorage_GetVoidPtr_                               = &c.FuncPrototype{Name: "ImGuiStorage_GetVoidPtr", OutType: c.Pointer, InTypes: _typs_PU32}
	_func_ImGuiStorage_GetVoidPtrRef_                            = &c.FuncPrototype{Name: "ImGuiStorage_GetVoidPtrRef", OutType: c.Pointer, InTypes: _typs_PU32P}
	_func_ImGuiStorage_SetAllInt_                                = &c.FuncPrototype{Name: "ImGuiStorage_SetAllInt", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImGuiStorage_SetBool_                                  = &c.FuncPrototype{Name: "ImGuiStorage_SetBool", OutType: c.Void, InTypes: _typs_PU32I32}
	_func_ImGuiStorage_SetFloat_                                 = &c.FuncPrototype{Name: "ImGuiStorage_SetFloat", OutType: c.Void, InTypes: _typs_PU32F32}
	_func_ImGuiStorage_SetInt_                                   = &c.FuncPrototype{Name: "ImGuiStorage_SetInt", OutType: c.Void, InTypes: _typs_PU32I32}
	_func_ImGuiStorage_SetVoidPtr_                               = &c.FuncPrototype{Name: "ImGuiStorage_SetVoidPtr", OutType: c.Void, InTypes: _typs_PU32P}
	_func_ImGuiStoragePair_destroy_                              = &c.FuncPrototype{Name: "ImGuiStoragePair_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiStoragePair_ImGuiStoragePair_Float_               = &c.FuncPrototype{Name: "ImGuiStoragePair_ImGuiStoragePair_Float", OutType: c.Pointer, InTypes: _typs_U32F32}
	_func_ImGuiStoragePair_ImGuiStoragePair_Int_                 = &c.FuncPrototype{Name: "ImGuiStoragePair_ImGuiStoragePair_Int", OutType: c.Pointer, InTypes: _typs_U32I32}
	_func_ImGuiStoragePair_ImGuiStoragePair_Ptr_                 = &c.FuncPrototype{Name: "ImGuiStoragePair_ImGuiStoragePair_Ptr", OutType: c.Pointer, InTypes: _typs_U32P}
	_func_ImGuiTextBuffer_append_                                = &c.FuncPrototype{Name: "ImGuiTextBuffer_append", OutType: c.Void, InTypes: _typs_PPP}
	_func_ImGuiTextBuffer_c_str_                                 = &c.FuncPrototype{Name: "ImGuiTextBuffer_c_str", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImGuiTextBuffer_clear_                                 = &c.FuncPrototype{Name: "ImGuiTextBuffer_clear", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiTextBuffer_empty_                                 = &c.FuncPrototype{Name: "ImGuiTextBuffer_empty", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiTextBuffer_end_                                   = &c.FuncPrototype{Name: "ImGuiTextBuffer_end", OutType: c.Pointer, InTypes: _typs_P}
	_func_ImGuiTextBuffer_reserve_                               = &c.FuncPrototype{Name: "ImGuiTextBuffer_reserve", OutType: c.Void, InTypes: _typs_PI32}
	_func_ImGuiTextBuffer_size_                                  = &c.FuncPrototype{Name: "ImGuiTextBuffer_size", OutType: c.I32, InTypes: _typs_P}
	_func_ImGuiViewport_destroy_                                 = &c.FuncPrototype{Name: "ImGuiViewport_destroy", OutType: c.Void, InTypes: _typs_P}
	_func_ImGuiViewport_GetCenter_                               = &c.FuncPrototype{Name: "ImGuiViewport_GetCenter", OutType: c.Void, InTypes: _typs_PP}
	_func_ImGuiViewport_GetWorkCenter_                           = &c.FuncPrototype{Name: "ImGuiViewport_GetWorkCenter", OutType: c.Void, InTypes: _typs_PP}
	_func_ImGuiViewport_ImGuiViewport_                           = &c.FuncPrototype{Name: "ImGuiViewport_ImGuiViewport", OutType: c.Pointer, InTypes: nil}
)
